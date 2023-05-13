package broker

import (
	"context"

	"github.com/MihasBel/data-bus-publisher/internal/rep"
	"github.com/rs/zerolog"
)

type Broker struct {
	cfg Config
	p   rep.Publisher
	cc  rep.ConsumerCreator
	log *zerolog.Logger
}

func New(cfg Config, log zerolog.Logger, p rep.Publisher, cc rep.ConsumerCreator) *Broker {
	return &Broker{
		cfg: cfg,
		log: &log,
		p:   p,
		cc:  cc,
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
