package servicetest

import (
	"countries-info/internal/services"
	"os"
	"sync"
	"testing"
)

func TestFetchCountryInfoConcurrently(t *testing.T) {

	os.Setenv("COUNTRY_API_URL", "https://restcountries.com/v3.1/name/{country_name}")

	country := "India"

	var wg sync.WaitGroup
	numRequests := 20

	for i := 0; i < numRequests; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()

			_, err := services.FetchCountryInfo(country)

			if err != nil {
				t.Errorf("Error fetching the country info: %v", err)
			}

		}()

	}

	wg.Wait()

}
