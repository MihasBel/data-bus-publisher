package app

import (
	"context"
	"testing"
	"time"

	"github.com/rs/zerolog"
)

func TestApp_Start(t *testing.T) {
	app := New(Config{StartTimeout: 1 * time.Second}, zerolog.Nop())

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := app.Start(ctx)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
