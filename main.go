package main

import (
	"github.com/qbantek/to-localhost/internal/config"
	"github.com/qbantek/to-localhost/internal/routes"
)

func main() {
	cfg := config.NewConfig()

	routes.SetupRouter().Run(":" + cfg.Port)
}
