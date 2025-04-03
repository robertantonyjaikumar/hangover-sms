package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sms/models"
	"sms/routes"
	"syscall"
	"time"

	"github.com/robertantonyjaikumar/hangover-common/config"
	"github.com/robertantonyjaikumar/hangover-common/logger"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

// @title           SMS API
// @version         1.0
// @description     School Management System.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Antony Raj, Jaikumar, Richard
// @contact.email  robertantonyjaikumar@gmail.com

// @host      localhost:8081
// @BasePath  /v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	models.MigrateDB()
	// models.SeedDB()
	// Create CORS middleware options
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Allow React frontend origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	router := routes.NewRouter()
	// Apply CORS middleware to the router (wrap it with the cors handler)
	handler := corsHandler.Handler(router)
	// Create a new HTTP server with custom configuration
	server := &http.Server{
		Addr:    ":" + config.CFG.V.GetString("server.port"),
		Handler: handler,
	}

	// Create a channel to listen for OS signals (e.g., SIGINT, SIGTERM)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Run the server in a goroutine
	go func() {
		// Start listening and serving requests
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// Log any error that causes the server to stop
			logger.Panic("Failed to start HTTP server", zap.Error(err))
		}
	}()

	// Log that the server has started successfully
	logger.Info("Server started")

	// Wait for a signal to shut down
	<-stop

	// Graceful shutdown: Wait for 5 seconds before forcefully shutting down
	logger.Info("Shutting down server gracefully...")

	// Setting a timeout for graceful shutdown (5 seconds in this case)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		// Log if the server fails to shut down gracefully
		logger.Panic("Server shutdown failed", zap.Error(err))
	}

	logger.Info("Server shut down successfully")
}
