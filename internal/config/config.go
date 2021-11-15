package config

import "os"

const defaultPort = "5000"

type Config struct {
	Port string
}

func NewConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	return &Config{port}
}
