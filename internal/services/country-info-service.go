package services

import (
	"countries-info/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func FetchCountryInfo(country string) (*models.CountryResponse, error) {
	baseURL := os.Getenv("COUNTRY_API_URL")

	if baseURL == "" {
		return nil, fmt.Errorf("COUNTRY_API_URL not set")
	}

	url := strings.Replace(
		baseURL,          // original URL
		"{country_name}", // text to replace
		country,          // replace with actual country passed by user
		1,                // replace only once
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call the API: %v", err)
	}
	defer resp.Body.Close()

	// fmt.Println("resp:", resp)

	// Handle non-200 responses gracefully
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d (%s)", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	// check for invalid country

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read API response: %w", err)
	}

	var raw []models.RestCountry

	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %v", err)
	}

	if len(raw) == 0 {
		return nil, fmt.Errorf("country not found: %s", country)
	}

	// fmt.Println("raw:", raw)

	apiResp := raw[0]

	// extract required fields from raw and map to CountryResponse struct
	// Handle missing capital gracefully
	capital := "N/A"
	if len(apiResp.Capital) > 0 {
		capital = apiResp.Capital[0]
	}
	
	currencySymbol := "N/A"

	for _, v := range apiResp.Currencies {
		currencySymbol = v.Symbol
		break // only need the first currency
	}

	return &models.CountryResponse{
		Name:       apiResp.Name.Common,
		Capital:    capital,
		Currency:   currencySymbol,
		Population: int64(apiResp.Population),
	}, nil

}
