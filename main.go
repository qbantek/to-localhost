package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/qbantek/to-localhost/internal/config"
	"github.com/qbantek/to-localhost/internal/routes"
)

const shutdownTimeout = 5 * time.Second

func main() {
	cfg := config.NewConfig()
	router := routes.NewEngine()

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	startServer(srv)
}

func startServer(srv *http.Server) {
	// Start the server and listen for errors
	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- srv.ListenAndServe()
	}()

	// Listen for an interrupt or terminate signal from the OS.
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	// Block until we receive a signal or the server returns an error.
	select {
	case err := <-serverErrors:
		log.Fatal(err)
		return
	case <-osSignals:
		log.Println("Stopping server...")
	}

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:" + err.Error())
	}

	log.Println("Server stopped")
}
