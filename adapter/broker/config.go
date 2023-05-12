package broker

import "time"

type Config struct {
	ServerURL       string        `env:"BROKER_SERVER"`
	AutoOffsetReset string        `env:"AUTO_OFFSET_RESET" envDefault:"earliest"`
	GroupTTL        time.Duration `env:"GROUP_TTL" envDefault:"60m"`
	GroupTimeoutMs  int           `env:"GROUP_TIMEOUT_MS" envDefault:"100"`
	MaxMessageBytes int           `env:"MAX_MESSAGE_BYTES" envDefault:"50000"`
}
