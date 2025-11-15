# country-info

A RESTful API built in Go to fetch comprehensive information about countries from the REST Countries API, with built-in caching for improved performance.

## Overview

`country-info` is a lightweight Go service that provides country information including name, capital, currency, and population. The API integrates with the [REST Countries](https://restcountries.com) external API and implements an in-memory caching layer to reduce API calls and improve response times.

## Features

- **Country Search**: Query country information by name
- **In-Memory Caching**: Caches country data to minimize external API calls
- **Health Check Endpoint**: Monitor server readiness
- **Error Handling**: Comprehensive error handling with meaningful error messages
- **Environment Configuration**: Configurable port via `.env` file
- **Graceful Shutdown**: Handles server shutdown signals gracefully

## Project Structure

```
country-info/
├── cmd/
│   └── country-info/
│       └── main.go              # Application entry point
├── internal/
│   ├── handlers/
│   │   ├── country-info-handler.go    # HTTP handler for country search
│   │   └── server-rediness.go         # Health check handler
│   ├── models/
│   │   ├── country-response-model.go  # Country response structure
│   │   └── rest-country-model.go      # REST Countries API response model
│   ├── routes/
│   │   └── routes.go            # API route definitions
│   ├── services/
│   │   ├── country-info-service.go    # Business logic for country lookup
│   │   └── servicehelpers/
│   │       └── helpers.go             # Service helper functions
│   └── utils/
│       ├── cache.go             # Cache interface
│       ├── inmemory-cache.go    # In-memory cache implementation
│       ├── json-response.go     # JSON response utilities
│       └── cache_test.go        # Cache tests
├── go.mod                        # Go module file
└── README.md                     # This file
```

## API Endpoints

### Health Check
```
GET /health
```
Returns server readiness status.

**Response (200 OK):**
```json
{
  "message": "Server is ready"
}
```

### Search Country Information
```
GET /api/countries/search?name=<country_name>
```
Retrieves country information by name.

**Query Parameters:**
- `name` (required): The name of the country to search for

**Response (200 OK):**
```json
{
  "name": "India",
  "capital": "New Delhi",
  "currency": "Indian Rupee",
  "population": 1417173173
}
```

**Error Response (400 Bad Request):**
```json
{
  "error": "Country name is required"
}
```

## Getting Started

### Prerequisites

- Go 1.24.2 or higher
- Internet connection (for REST Countries API integration)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/vinit-jpl/country-info.git
cd country-info
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the project root (optional):
```
PORT=8080
```

If not specified, the server defaults to port `8080`.

### Running the Application

```bash
go run cmd/country-info/main.go
```

The server will start and listen on the configured port (default: 8080).

### Example Usage

```bash
# Health check
curl http://localhost:8080/health

# Search for a country
curl "http://localhost:8080/api/countries/search?name=India"

# Search for another country
curl "http://localhost:8080/api/countries/search?name=France"
```

## Architecture

### Request Flow

1. **HTTP Handler** (`country-info-handler.go`): Receives incoming requests and validates input
2. **Service Layer** (`country-info-service.go`): Orchestrates business logic
   - Validates country input
   - Checks in-memory cache
   - Fetches from REST Countries API if not cached
   - Parses and transforms API response
   - Caches result for future requests
3. **Models**: Define data structures for API responses
4. **Utils**: Provide caching and response utilities

### Caching Strategy

The application implements an in-memory cache to store country information:
- First request for a country fetches from REST Countries API
- Subsequent requests return cached data (faster response)
- Reduces external API calls and improves performance

## Dependencies

- `github.com/joho/godotenv v1.5.1` - For loading environment variables from `.env` file

## Testing

Run tests with:
```bash
go test ./...
```

Tests are included for the caching mechanism in `internal/utils/cache_test.go`.

## Development

### Adding New Fields

To add new country fields:

1. Update the `CountryResponse` model in `internal/models/country-response-model.go`
2. Update the mapping logic in `internal/services/servicehelpers/helpers.go`
3. Update the `RestCountryModel` if adding external API fields

### Extending Routes

Add new routes in `internal/routes/routes.go` and corresponding handlers in `internal/handlers/`.

## Error Handling

The API provides descriptive error messages:
- **Missing query parameter**: "Country name is required"
- **Invalid country input**: "invalid country name: {name}"
- **API errors**: Forwarded from REST Countries API

## Future Enhancements

- Add pagination for large result sets
- Implement persistent caching (Redis)
- Add rate limiting
- Support multiple country search
- Add filtering by region or language
- Implement user authentication

## License

This project is open source and available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues for bugs and feature requests.

## Author

- **vinit-jpl** - [GitHub Profile](https://github.com/vinit-jpl)
