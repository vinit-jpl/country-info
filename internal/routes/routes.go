package routes

import (
	"countries-info/internal/handlers"
	"net/http"
)

func InitializeRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.HandlerReadiness)

	return mux
}
