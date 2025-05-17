package main

import (
	"errors"
	"net/http"
	"strings"
)

//GetAPIKey extracts an API key from the headers
// of an http request

// Example
// Authentication: ApiKey {insert apiKey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("No authentication info found")

	}

	vals := strings.Split(val, "")
	if len(vals) != 2 {
		return "", errors.New("Malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed apikey error auth header")

	}
	return vals[1], nil
}
