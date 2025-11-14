package services

import (
	"countries-info/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// httpClient is used to make API requests with a fixed timeout.
// This prevents the application from waiting forever if the API is slow.
var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

// Build URL
func buildURL(country string) (string, error) {

	baseURL := os.Getenv("COUNTRY_API_URL")

	if baseURL == "" {
		return "", fmt.Errorf("COUNTRY_API_URL not set")
	}

	url := strings.Replace(
		baseURL,          // original URL
		"{country_name}", // text to replace
		country,          // replace with actual country passed by user
		1,                // replace only once
	)

	return url, nil

}

// Call API
func callAPI(url string) ([]byte, error) {
	resp, err := httpClient.Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to call the API: %v", err)
	}

	defer resp.Body.Close()

	// Handle non-200 responses gracefully
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d (%s)", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("Failed to read API response: %w", err)
	}

	return body, nil

}

// Parse API response
func parseAPIResponse(body []byte) (*models.RestCountry, error) {
	var raw []models.RestCountry

	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}

	if len(raw) == 0 {
		return nil, fmt.Errorf("empty response from API or country not found")
	}

	return &raw[0], nil

}

// Extract required Fields
func extractCapital(data *models.RestCountry) string {

	if len(data.Capital) > 0 {
		return data.Capital[0]
	}
	return "N/A"
}

func extractCurrency(data *models.RestCountry) string {

	for _, v := range data.Currencies {
		return v.Symbol // return first found
	}
	return "N/A"
}

// Build country response
func mapToCountryResponse(apiResp *models.RestCountry) *models.CountryResponse {
	return &models.CountryResponse{
		Name:       apiResp.Name.Common,
		Capital:    extractCapital(apiResp),
		Currency:   extractCurrency(apiResp),
		Population: int64(apiResp.Population),
	}
}

// Call helpers here
func FetchCountryInfo(country string) (*models.CountryResponse, error) {
	url, err := buildURL(country)

	if err != nil {
		return nil, err
	}

	body, err := callAPI(url)

	if err != nil {
		return nil, err
	}

	apiResp, err := parseAPIResponse(body)

	if err != nil {
		return nil, err
	}

	return mapToCountryResponse(apiResp), nil
}
