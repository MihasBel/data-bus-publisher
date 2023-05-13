package rep

import (
	"context"

	"github.com/MihasBel/data-bus-publisher/internal/models"
)

type SubscriptionManager interface {
	Subscribe(ctx context.Context, subscriber *models.Subscriber) error
	Unsubscribe(ctx context.Context, id string, msgType string) error
}
