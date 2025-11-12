package main

import (
	"context"
	"countries-info/internal/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func loadEnv() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func setupServer() (*http.Server, *http.ServeMux) {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080" // Default port if not specified
	}

	mux := routes.InitializeRoutes()

	// return a new HTTP server instance
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	return srv, mux
}

func startServer(srv *http.Server) {
	go func() {
		log.Println("Server started on port", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe failed: %v", err)
		}
	}()
}

func gracefulShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}

func main() {

	loadEnv()

	srv, _ := setupServer()

	startServer(srv)

	gracefulShutdown(srv)
}
