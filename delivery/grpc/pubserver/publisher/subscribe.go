package publisher

import (
	"context"

	"github.com/MihasBel/data-bus-publisher/delivery/grpc/gen/v1/publisher"
	"github.com/MihasBel/data-bus-publisher/internal/models"
)

func (s *Server) Subscribe(request *publisher.SubscriptionRequest, stream publisher.PubSubService_SubscribeServer) error {
	ctx, cancel := context.WithCancel(stream.Context())
	subscriber := &models.Subscriber{ // TODO auth JWT token to get ID
		ID:          request.SubscriberId,
		Stream:      stream,
		MessageType: request.MessageType,
		Cancel:      cancel,
	}
	if err := s.m.Subscribe(ctx, subscriber); err != nil {
		return err
	}

	<-stream.Context().Done()

	return stream.Context().Err()
}
