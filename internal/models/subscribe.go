package models

import (
	"context"
	"github.com/MihasBel/data-bus-publisher/delivery/grpc/gen/v1/publisher"
	"github.com/pkg/errors"
)

type Subscriber struct {
	ID          string
	Stream      publisher.PubSubService_SubscribeServer
	MessageType string
	Cancel      context.CancelFunc
}

func (s *Subscriber) IsValid() error {
	if s.ID == "" {
		return errors.New("subscriber ID cannot be empty")
	}
	if s.Stream == nil {
		return errors.New("stream cannot be nil")
	}
	if s.MessageType == "" {
		return errors.New("message type cannot be empty")
	}
	if s.Cancel == nil {
		return errors.New("CancelFunc ch cannot be nil")
	}
	return nil
}
