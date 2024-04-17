package cocoon

import (
	"os"
)

type Config struct {
  AmqpUrl string
  DatabaseUrl string
}

func NewConfig() (*Config, error) {
  config := &Config{
    AmqpUrl: os.Getenv("DATABASE_URL"),
    DatabaseUrl: os.Getenv("DATABASE_URL"),
  }

	return config, nil
}
