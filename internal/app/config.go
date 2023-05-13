package app

import (
	"time"

	"github.com/MihasBel/data-bus-publisher/adapter/broker"
	"github.com/MihasBel/data-bus-publisher/delivery/grpc/pubserver"
)

type Config struct {
	LogLevel     string `env:"LOG_LEVEL" envDefault:"info"`
	BrokerConfig broker.Config
	GRPCConfig   pubserver.Config

	StartTimeout time.Duration `env:"START_TIMEOUT"`
	StopTimeout  time.Duration `env:"STOP_TIMEOUT"`
}
