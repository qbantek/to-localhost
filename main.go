package main

import (
	"os"

	"github.com/qbantek/to-localhost/internal/routes"
)

const defaultPort = "5000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	routes.SetupRouter().Run(":" + port)
}
