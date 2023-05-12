package subscription

import (
	"context"
	"fmt"
	"github.com/MihasBel/data-bus-publisher/delivery/grpc/gen/v1/publisher"
	"github.com/MihasBel/data-bus-publisher/internal/models"
	"github.com/MihasBel/data-bus/broker/model"
	"github.com/pkg/errors"
	"sync"
)

type Service struct {
	subMap map[string]*models.Subscriber
	mu     sync.Mutex
}

func New() *Service {
	return &Service{
		subMap: make(map[string]*models.Subscriber),
	}
}

func (s *Service) Subscribe(subscriber *models.Subscriber) error {
	err := subscriber.IsValid()
	if err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	v, ok := s.subMap[subscriber.ID]
	if ok {
		v.Stream = subscriber.Stream
		return nil
	}
	s.subMap[subscriber.ID] = subscriber
	return nil
}

func (s *Service) Unsubscribe(_ context.Context, subscriberID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	v, ok := s.subMap[subscriberID]
	if ok {
		v.Unsubscribe <- struct{}{}
		delete(s.subMap, subscriberID)
		return nil
	}
	return nil
}

func (s *Service) Get(subscriberID string) (*models.Subscriber, error) {
	sub, ok := s.subMap[subscriberID]
	if !ok {
		return nil, fmt.Errorf("subscriber with ID %s not found", subscriberID)
	}
	return sub, nil
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

// Start starts Service.
func (s *Service) Start(_ context.Context) error {
	return nil
}

// Stop stops Service.
func (s *Service) Stop(_ context.Context) error {
	return nil
}
