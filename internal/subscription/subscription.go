package subscription

import (
	"context"
	"fmt"
	"github.com/MihasBel/data-bus-publisher/adapter/broker"
	"github.com/MihasBel/data-bus-publisher/internal/models"
	"sync"
)

const (
	subkeyf = "id_%s_type_%s"
)

type Service struct {
	subMap map[string]*models.Subscriber
	b      *broker.Broker
	mu     sync.Mutex
}

func New(b *broker.Broker) *Service {
	return &Service{
		b:      b,
		subMap: make(map[string]*models.Subscriber),
	}
}

func (s *Service) Subscribe(ctx context.Context, subscriber *models.Subscriber) error {
	err := subscriber.IsValid()
	if err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	key := fmt.Sprintf(subkeyf, subscriber.ID, subscriber.MessageType)
	v, ok := s.subMap[key]
	if ok {
		v.Cancel()
		v.Stream = subscriber.Stream
	} else {
		s.subMap[key] = subscriber
	}
	return s.b.HandleConsumer(ctx, subscriber)
}

func (s *Service) Unsubscribe(_ context.Context, subscriberID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	v, ok := s.subMap[subscriberID]
	if ok {
		v.Cancel()
		delete(s.subMap, subscriberID)
		return nil
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
