package app

import (
	"context"
	"github.com/MihasBel/data-bus-publisher/adapter/broker"
	grpcServ "github.com/MihasBel/data-bus-publisher/delivery/grpc/pubserver"
	"github.com/MihasBel/data-bus-publisher/internal/subscription"
	"github.com/MihasBel/data-bus-publisher/pkg/lifecycle"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"time"
)

const (
	FmtCannotStart = "cannot start %q"
)

var (
	ErrStartTimeout    = errors.New("start timeout")
	ErrShutdownTimeout = errors.New("shutdown timeout")
)

type (
	App struct {
		log  *zerolog.Logger
		cmps []cmp
		cfg  Config
	}
	cmp struct {
		Service lifecycle.Lifecycle
		Name    string
	}
)

func New(cfg Config, l zerolog.Logger) *App {
	l = l.With().Str("cmp", "app").Logger()

	return &App{
		log: &l,
		cfg: cfg,
	}
}

func (a *App) Start(ctx context.Context) error {
	a.log.Info().Msg("starting app")
	manager := subscription.New()
	b := broker.New(a.cfg.BrokerConfig, *a.log, manager, manager)
	grpcServer := grpcServ.New(a.cfg.GRPCConfig, *a.log, manager, b)

	a.cmps = append(
		a.cmps,
		cmp{manager, "subManager"},
		cmp{b, "broker"},
		cmp{grpcServer, "grpcServer"},
	)

	okCh, errCh := make(chan struct{}), make(chan error)

	go func() {
		for _, c := range a.cmps {
			a.log.Info().Msgf("%v is starting", c.Name)

			if err := c.Service.Start(ctx); err != nil {
				a.log.Error().Err(err).Msgf(FmtCannotStart, c.Name)
				errCh <- errors.Wrapf(err, FmtCannotStart, c.Name)

				return
			}

			a.log.Info().Msgf("%v started", c.Name)
		}
		okCh <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return ErrStartTimeout
	case err := <-errCh:
		return err
	case <-okCh:
		a.log.Info().Msg("Application started!")
		return nil
	}
}

func (a *App) Stop(ctx context.Context) error {
	a.log.Info().Msg("shutting down service...")

	okCh, errCh := make(chan struct{}), make(chan error)

	go func() {
		for i := len(a.cmps) - 1; i > 0; i-- {
			c := a.cmps[i]
			a.log.Info().Msgf("stopping %q...", c.Name)

			if err := c.Service.Stop(ctx); err != nil {
				a.log.Error().Err(err).Msgf("cannot stop %q", c.Name)
				errCh <- err

				return
			}
		}

		okCh <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return ErrShutdownTimeout
	case err := <-errCh:
		return err
	case <-okCh:
		a.log.Info().Msg("Application stopped!")
		return nil
	}
}

func (a *App) GetStartTimeout() time.Duration { return a.cfg.StartTimeout }
func (a *App) GetStopTimeout() time.Duration  { return a.cfg.StopTimeout }
