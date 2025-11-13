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

	url := strings.Replace(
		baseURL,          // original URL
		"{country_name}", // text to replace
		country,          // replace with actual country passed by user
		1,                // replace only once
	)

	resp, err := http.Get(url)

	// fmt.Println("resp:", resp)

	if err != nil {
		return nil, fmt.Errorf("Failed to call the API: %v", err)
	}

	defer resp.Body.Close()

	// check for invalid country

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("Invalid Country: %s", country)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)

	var data models.CountryResponse

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	// fmt.Printf("Fetched data: %+v\n", data)

	return &data, nil

}
