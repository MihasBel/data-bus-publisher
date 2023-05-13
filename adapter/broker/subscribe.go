package broker

import (
	"context"
	"fmt"

	"github.com/MihasBel/data-bus-publisher/internal/models"
	"github.com/MihasBel/data-bus/broker/bustopic"
	"github.com/MihasBel/data-bus/broker/model"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

func (b *Broker) HandleConsumer(ctx context.Context, subscriber *models.Subscriber) error {
	if !bustopic.IsInTopics(subscriber.MessageType) {
		return errors.Errorf("Unsupported msg type: %s", subscriber.MessageType)
	}
	c, err := b.cc.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":          b.cfg.ServerURL,
		"group.id":                   subscriber.ID,
		"auto.offset.reset":          b.cfg.AutoOffsetReset,
		"allow.auto.create.topics":   true,
		"enable.auto.offset.store":   false,
		"queued.max.messages.kbytes": b.cfg.MaxMessageBytes,
	})
	if err != nil {
		return fmt.Errorf("failed to create consumer: %w", err)
	}

	err = c.Subscribe(subscriber.MessageType, nil)
	if err != nil {
		return fmt.Errorf("failed to subscribe to topic: %w", err)
	}

	go b.consume(ctx, subscriber, c)

	return nil
}

func (b *Broker) consume(ctx context.Context, sub *models.Subscriber, c *kafka.Consumer) {
	defer func(c *kafka.Consumer) {
		err := c.Close()
		if err != nil {
			b.log.Error().Err(err)
		}
	}(c)
	for {
		var err error
		select {
		case <-ctx.Done():
			b.log.Info().Msgf("cancel ctx in HandleConsumer go func id:%s", sub.ID)
			return
		default:
			ev := c.Poll(b.cfg.GroupTimeoutMs)
			if ev == nil {
				continue
			}

			switch msg := ev.(type) {
			case *kafka.Message:
				if err := sub.IsValid(); err != nil {
					b.log.Error().Err(err).Msgf("Get subscriber ID:%v", sub.ID)
					continue
				}

				err = b.p.Publish(ctx, sub, &model.Message{
					MsgType: sub.MessageType,
					Data:    msg.Value,
				})
				if err != nil {
					b.log.Error().Err(err).Msgf("Publish to subscriber ID:%v", sub.ID)
					continue
				}
				_, err = c.StoreOffsets(
					[]kafka.TopicPartition{
						{Topic: msg.TopicPartition.Topic,
							Partition: msg.TopicPartition.Partition,
							Offset:    msg.TopicPartition.Offset + 1},
					})

				if err != nil {
					b.log.Error().Err(err).Msgf("StoreOffsets subscriber ID:%v", sub.ID)
					continue
				}
			case kafka.Error:
				b.log.Error().Msgf("kafka.Error in subscriber ID:%v error:%s", sub.ID, msg.Error())
			default:
			}
		}
	}
}
