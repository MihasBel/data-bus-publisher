package app

import (
	"github.com/MihasBel/data-bus-publisher/delivery/grpc/server"
	"time"
)

type Config struct {
	LogLevel string   `env:"LOG_LEVEL" envDefault:"info"`
	MsgTypes []string `env:"MSG_TYPES" required:"true"`
	//KafkaConfig broker.Config
	GRPCConfig server.Config

	StartTimeout time.Duration `env:"START_TIMEOUT"`
	StopTimeout  time.Duration `env:"STOP_TIMEOUT"`
}
