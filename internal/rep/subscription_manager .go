package rep

import (
	"context"
	"github.com/MihasBel/data-bus-publisher/internal/models"
)

type SubscriptionManager interface {
	Subscribe(subscriber *models.Subscriber) error
	Unsubscribe(ctx context.Context, id string) error
	Get(subscriberID string) (*models.Subscriber, error)
}
