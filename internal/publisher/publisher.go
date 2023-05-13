package publisher

import (
	"context"
	"github.com/MihasBel/data-bus-publisher/delivery/grpc/gen/v1/publisher"
	"github.com/MihasBel/data-bus-publisher/internal/models"
	"github.com/MihasBel/data-bus/broker/model"
	"github.com/pkg/errors"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Publish(_ context.Context, subscriber *models.Subscriber, msg *model.Message) error {
	err := subscriber.IsValid()
	if err != nil {
		return err
	}
	grpcMsg := &publisher.Message{
		Type: msg.MsgType,
		Data: msg.Data,
	}

	if err := subscriber.Stream.Send(grpcMsg); err != nil {
		return errors.Wrapf(err, "failed to send message to subscriber %s", subscriber.ID)
	}

	return nil
}
