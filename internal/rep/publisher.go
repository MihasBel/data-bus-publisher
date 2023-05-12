package rep

import (
	"context"
	"github.com/MihasBel/data-bus-publisher/internal/models"
	"github.com/MihasBel/data-bus/broker/model"
)

type Publisher interface {
	Publish(ctx context.Context, sub *models.Subscriber, msg *model.Message) error
}
