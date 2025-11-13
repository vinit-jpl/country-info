package models

// JSON response from api starts from [, so we need to decode it in slice of objects.
type CountryResponse []struct {

	// name field is itself an object so we use nested struct
	Name struct {
		Common string `json:"common"`
	} `json:"name"`

	// capital gives array of strings
	Capital []string `json:"capital"`

	// gives simple integer
	Population int `json:"population"`

	// key value map for currencies
	// key => Currency Code, like INR
	// value => another struct with name and symbol
	Currencies map[string]CurrencyObj `json:"currencies"`
}

type CurrencyObj struct {
	Name string `json:"name"`

	Symbol string `json:"symbol"`
}

/* response from api
[
	{
		"name": {
			"common": India
		},

		"capital": [
			"New Delhi"
		],

		"population": 1417492000,

		"currencies": {
			"INR": {
				"symbol": "₹",
				"name": "Indian rupee"
			}
		},



	}
]
*/
