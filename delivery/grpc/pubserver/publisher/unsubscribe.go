package publisher

import (
	"context"

	"github.com/MihasBel/data-bus-publisher/delivery/grpc/gen/v1/publisher"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Unsubscribe(ctx context.Context, request *publisher.SubscriptionRequest) (*emptypb.Empty, error) {
	err := s.m.Unsubscribe(ctx, request.SubscriberId, request.MessageType)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
