package mocks

import (
	"context"
	"github.com/MihasBel/data-bus-publisher/internal/models"
	"github.com/stretchr/testify/mock"
)

type BrokerMock struct {
	mock.Mock
}

func (b *BrokerMock) HandleConsumer(ctx context.Context, subscriber *models.Subscriber) error {
	args := b.Called(ctx, subscriber)
	return args.Error(0)
}
