package publisher

import (
	"github.com/MihasBel/data-bus-publisher/delivery/grpc/gen/v1/publisher"
	"github.com/MihasBel/data-bus-publisher/internal/rep"
)

// Server SubscriptionManager
type Server struct {
	publisher.UnimplementedPubSubServiceServer
	m rep.SubscriptionManager
}

// New constructor
func New(m rep.SubscriptionManager) *Server {
	return &Server{
		m: m,
	}
}
