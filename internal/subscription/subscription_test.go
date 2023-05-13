package subscription

import (
	"context"
	"github.com/MihasBel/data-bus-publisher/mocks"
	"github.com/pkg/errors"
	"testing"

	"github.com/MihasBel/data-bus-publisher/internal/models"
	"github.com/stretchr/testify/mock"
)

func TestService_Subscribe(t *testing.T) {
	tests := []struct {
		name        string
		subscriber  *models.Subscriber
		wantErr     error
		brokerError error
	}{
		{
			name: "valid subscriber, not yet subscribed",
			subscriber: &models.Subscriber{
				ID:          "123",
				MessageType: "type1",
			},
			wantErr: nil,
		},
		{
			name: "valid subscriber, already subscribed",
			subscriber: &models.Subscriber{
				ID:          "123",
				MessageType: "type1",
			},
			wantErr: nil,
		},
		{
			name: "invalid subscriber",
			subscriber: &models.Subscriber{
				ID:          "",
				MessageType: "type1",
			},
			wantErr: errors.New("subscriber ID cannot be empty"),
		},
		{
			name: "broker error",
			subscriber: &models.Subscriber{
				ID:          "123",
				MessageType: "type1",
			},
			wantErr:     errors.New("broker error"),
			brokerError: errors.New("broker error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			broker := mocks.BrokerMock{}
			broker.On("HandleConsumer", mock.Anything, tt.subscriber).Return(tt.brokerError)

			service := New(&broker)

			err := service.Subscribe(context.Background(), tt.subscriber)

			if err != tt.wantErr {
				t.Errorf("Service.Subscribe() error = %v, wantErr %v", err, tt.wantErr)
			}

			broker.AssertExpectations(t)
		})
	}
}
