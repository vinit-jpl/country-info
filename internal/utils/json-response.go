package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// payload is typed as interface{} so the function can accept any Go value
// (structs, maps, slices, strings, etc.) and encode it into JSON.
// This makes RespondWithJSON reusable for all response types instead of
// restricting it to a single data structure.

func RespondWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {

	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Error in marshalling the payload: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}
