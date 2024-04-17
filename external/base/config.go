package cocoon_connect

import (
	"os"
)

type Config struct {
    AmqpURL string
}

func NewConfig() (*Config, error) {
  config := &Config{
    AmqpURL: os.Getenv("AMQP_URL"),
  }

  return config, nil
}
