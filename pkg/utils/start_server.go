package utils

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func StartServerWithGracefulShutdown(e *echo.Echo, url string) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			log.Panic("Failed to shut down the server - Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := e.Start(url); err != nil {
		log.Fatal("Failed to start the server :- Reason: %v", err)
	}

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func StartServer(e *echo.Echo, url string) {

	// Run server.
	if err := e.Start(url); err != nil {
		log.Fatal("Failed to shut down the server - Reason: %v", err)
	}
}
