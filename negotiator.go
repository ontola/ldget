package main

import (
	"net/http"
)

// Negotiator -- Tries to fetch a resource using HTTP content negotiation
func Negotiator(url string) (*http.Response, error) {

	resp, err := http.Get(url)

	return resp, err
}
