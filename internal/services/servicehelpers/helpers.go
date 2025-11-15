package servicehelpers

import (
	"countries-info/internal/models"
	"countries-info/internal/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var CountryInfoCache = utils.NewCache()

// httpClient is used to make API requests with a fixed timeout.
// This prevents the application from waiting forever if the API is slow.
var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

// Validate country input
func IsValidCountryInput(country string) bool {
	country = strings.TrimSpace(country)

	// Reject empty string
	if country == "" {
		return false
	}

	// Reject strings with no alphabet characters
	hasLetter := false
	for _, ch := range country {
		if ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') {
			hasLetter = true
			break
		}
	}

	return hasLetter
}

// Build URL
func BuildURL(country string) (string, error) {

	if !IsValidCountryInput(country) {
		return "", fmt.Errorf("invalid country name: %s", country)
	}

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
func CallAPI(url string) ([]byte, error) {
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
		return nil, fmt.Errorf("failed to read API response: %w", err)
	}

	return body, nil

}

// Parse API response
func ParseAPIResponse(body []byte) (*models.RestCountry, error) {
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
func MapToCountryResponse(apiResp *models.RestCountry) *models.CountryResponse {
	return &models.CountryResponse{
		Name:       apiResp.Name.Common,
		Capital:    extractCapital(apiResp),
		Currency:   extractCurrency(apiResp),
		Population: int64(apiResp.Population),
	}
}

func FetchCountryFromCache(country string) (*models.CountryResponse, bool) {

	// validate country input
	if !IsValidCountryInput(country) {
		log.Printf("Invalid country input: '%s', skipping cache lookup", country)
	}
	cacheKey := strings.ToLower(country)

	cachedData, found := CountryInfoCache.Get(cacheKey)

	if !found {
		log.Printf("Country %s not found in cache, calling the 3rd party API", country)
		return nil, false
	}

	var cachedResp models.CountryResponse

	if err := json.Unmarshal(cachedData.([]byte), &cachedResp); err != nil {
		return nil, false
	}

	return &cachedResp, true
}

func SetCountryInCache(country string, resp *models.CountryResponse) {

	if !IsValidCountryInput(country) {
		log.Printf("Invalid country '%s'. Skipping cache save.", country)
		return
	}

	cacheKey := strings.ToLower(country)

	jsonData, err := json.Marshal(resp)

	if err != nil {
		return
	}

	CountryInfoCache.Set(cacheKey, jsonData)
	log.Printf("Saved country '%s' into cache\n", cacheKey)

}
