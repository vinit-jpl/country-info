package handlers

import (
	"countries-info/internal/utils"
	"net/http"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {

	utils.RespondWithJson(w, http.StatusOK, map[string]string{
		"message": "Server is ready",
	})
}
