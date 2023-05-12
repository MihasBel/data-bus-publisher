package pubserver

type Config struct {
	Host             string `env:"GRPC_URL" envDefault:"http://localhost:9090"`
	SecureConnection bool   `env:"GRPC_SECURE_CONNECTION" envDefault:"false"`
	MetricsEnabled   bool   `env:"METRICS_ENABLED" envDefault:"false"`
	StartTimeout     int    `env:"START_TIMEOUT" envDefault:"10"`
	StopTimeout      int    `env:"STOP_TIMEOUT" envDefault:"10"`
}
