package routes

import (
	"countries-info/internal/handlers"
	"net/http"
)

func InitializeRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.HandlerReadiness)

	mux.HandleFunc("GET /api/countries/search", handlers.GetCountryInfoHandler)

	return mux
}
