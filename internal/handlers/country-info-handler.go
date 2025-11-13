package handlers

import (
	"countries-info/internal/services"
	"countries-info/internal/utils"
	"net/http"
	"strings"
)

func GetCountryInfoHandler(w http.ResponseWriter, r *http.Request) {

	// extract country name from URL path

	country := r.URL.Query().Get("name")

	if strings.TrimSpace(country) == "" {
		utils.RespondWithJson(w, http.StatusBadRequest, map[string]string{
			"error": "Country name is required",
		})
		return
	}

	// else call the service to fetch country info

	result, err := services.FetchCountryInfo(country)

	if err != nil {
		utils.RespondWithJson(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	// successful response
	utils.RespondWithJson(w, http.StatusOK, result)

}
