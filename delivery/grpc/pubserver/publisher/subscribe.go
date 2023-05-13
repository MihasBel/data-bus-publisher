package publisher

import (
	"github.com/MihasBel/data-bus-publisher/delivery/grpc/gen/v1/publisher"
	"github.com/MihasBel/data-bus-publisher/internal/models"
)

func (s *Server) Subscribe(request *publisher.SubscriptionRequest, stream publisher.PubSubService_SubscribeServer) error {
	subscriber := &models.Subscriber{
		ID:          request.SubscriberId,
		Stream:      stream,
		MessageType: request.MessageType,
		Unsubscribe: make(chan struct{}),
	}
	if err := s.m.Subscribe(subscriber); err != nil {
		return err
	}
	if err := s.b.HandleConsumer(stream.Context(), subscriber); err != nil {
		return err
	}
	<-stream.Context().Done()

	return stream.Context().Err()
}
