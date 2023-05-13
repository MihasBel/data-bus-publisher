package broker

import "github.com/confluentinc/confluent-kafka-go/kafka"

type ConsCreator struct{}

func (r *ConsCreator) NewConsumer(config *kafka.ConfigMap) (*kafka.Consumer, error) {
	return kafka.NewConsumer(config)
}
