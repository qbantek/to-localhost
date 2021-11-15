package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/qbantek/to-localhost/internal/config"
	"github.com/qbantek/to-localhost/internal/routes"
)

func main() {
	cfg := config.NewConfig()
	router := routes.NewEngine()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}

	// Create a channel to listen for server errors
	// Start the server and listen for non-HTTP errors
	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- srv.ListenAndServe()
	}()

	// Create a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	// Block until we receive our signal.
	select {
	case err := <-serverErrors:
		log.Fatal(err)
		return
	case <-osSignals:
		log.Println("Shutting down server...")
	}

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:" + err.Error())
	}

	log.Println("Server stopped")
}
