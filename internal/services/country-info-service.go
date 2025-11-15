package services

import (
	"countries-info/internal/models"
	"countries-info/internal/services/servicehelpers"
	"fmt"
)

// Call helpers here
func FetchCountryInfo(country string) (*models.CountryResponse, error) {

	// Validate input
	if !servicehelpers.IsValidCountryInput(country) {
		return nil, fmt.Errorf("invalid country name: %s", country)
	}

	// Check cache first
	if cached, ok := servicehelpers.FetchCountryFromCache(country); ok {
		return cached, nil
	}

	// Not in cache, proceed to call API

	// build URL
	url, err := servicehelpers.BuildURL(country)

	if err != nil {
		return nil, err
	}

	// call API
	body, err := servicehelpers.CallAPI(url)

	if err != nil {
		return nil, err
	}

	// parse API response
	apiResp, err := servicehelpers.ParseAPIResponse(body)

	if err != nil {
		return nil, err
	}

	// map API -> final response
	finalResp := servicehelpers.MapToCountryResponse(apiResp)

	// Store in cache
	servicehelpers.SetCountryInCache(country, finalResp)

	return finalResp, nil
}
