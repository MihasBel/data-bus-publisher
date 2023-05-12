package broker

import (
	"context"
	"github.com/rs/zerolog"

	"github.com/MihasBel/data-bus-publisher/internal/rep"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Broker struct {
	cfg       Config
	m         rep.SubscriptionManager
	p         rep.Publisher
	log       *zerolog.Logger
	consumers []*kafka.Consumer
}

func New(cfg Config, log zerolog.Logger, m rep.SubscriptionManager, p rep.Publisher) *Broker {
	return &Broker{
		cfg:       cfg,
		log:       &log,
		m:         m,
		p:         p,
		consumers: make([]*kafka.Consumer, 0),
	}
}

// Start starts broker.
func (b *Broker) Start(_ context.Context) error {
	return nil
}

// Stop stops broker.
func (b *Broker) Stop(_ context.Context) error {
	return nil
}
