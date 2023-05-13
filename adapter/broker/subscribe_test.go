package broker

import (
	"context"
	"errors"
	"testing"

	"github.com/MihasBel/data-bus-publisher/internal/models"
	"github.com/MihasBel/data-bus-publisher/mocks"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestBroker_HandleConsumer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConsumerCreator := mocks.NewMockConsumerCreator(ctrl)
	mockPublisher := mocks.NewMockPublisher(ctrl)

	subscriber := &models.Subscriber{
		ID:          "test_subscriber",
		MessageType: "info",
	}

	b := &Broker{
		cfg: Config{},
		p:   mockPublisher,
		log: &zerolog.Logger{},
		cc:  mockConsumerCreator,
	}
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":          b.cfg.ServerURL,
		"group.id":                   subscriber.ID,
		"auto.offset.reset":          b.cfg.AutoOffsetReset,
		"allow.auto.create.topics":   true,
		"enable.auto.offset.store":   false,
		"queued.max.messages.kbytes": b.cfg.MaxMessageBytes,
	}

	/*t.Run("Success", func(t *testing.T) {
		mockConsumerCreator.EXPECT().NewConsumer(configMap).Return(&kafka.Consumer{}, nil)
		err := b.HandleConsumer(context.Background(), subscriber)
		assert.NoError(t, err)
	})*/

	t.Run("NewConsumerError", func(t *testing.T) {
		mockConsumerCreator.EXPECT().NewConsumer(configMap).Return(nil, errors.New("some error"))
		err := b.HandleConsumer(context.Background(), subscriber)
		assert.Error(t, err)
	})
}
