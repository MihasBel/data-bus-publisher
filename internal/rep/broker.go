package rep

import (
	"context"

	"github.com/MihasBel/data-bus-publisher/internal/models"
)

type Broker interface {
	HandleConsumer(ctx context.Context, subscriber *models.Subscriber) error
}
