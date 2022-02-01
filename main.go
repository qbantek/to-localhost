package main

import (
	"log"

	"github.com/qbantek/to-localhost/internal/config"
	"github.com/qbantek/to-localhost/internal/server"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	server.Start(cfg)
}
