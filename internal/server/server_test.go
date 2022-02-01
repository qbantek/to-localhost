package server_test

import (
	"testing"

	"github.com/qbantek/to-localhost/internal/config"
	"github.com/qbantek/to-localhost/internal/server"
)

// test server.Start()
func TestServerStart(t *testing.T) {
	server.Start(&config.Config{})
}
