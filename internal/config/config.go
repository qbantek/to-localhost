package config

import (
	"fmt"
	"os"

	"github.com/qbantek/to-localhost/internal/port"
)

const defaultPort = "5000"

type Config struct {
	Port string
}

func NewConfig() (*Config, error) {
	p := os.Getenv("PORT")
	if p == "" {
		p = defaultPort
	}

	// validate port
	_, err := port.NewPort(p)
	if err != nil {
		return nil, fmt.Errorf("NewConfig: %s", err)
	}

	return &Config{p}, nil
}
