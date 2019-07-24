package main

import (
	"net/http"

	"github.com/knakk/rdf"

	// "strings"d
	"log"
	"regexp"

	rdfmimetypes "github.com/ontola/ldget/rdfmimetypes"
)

// Negotiator -- Tries to fetch a resource using HTTP content negotiation
func Negotiator(url string) (*http.Response, rdf.Format, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Accept", acceptString())
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("HTTP status code: %v\n", resp.StatusCode)
	}
	respType := resp.Header.Get("Content-Type")
	if respType == "" {
		log.Fatal("No Content-Type header present on response!")
	}
	format := findFormat(respType)

	return resp, format, err
}

var acceptSelector, _ = regexp.Compile(`(.*);`)

func findFormat(header string) rdf.Format {
	matches := acceptSelector.FindStringSubmatch(header)
	headerFixed := header
	if len(matches) > 0 {
		headerFixed = matches[1]
	}
	for _, mapping := range rdfmimetypes.ContentTypes {
		if mapping.Header == headerFixed {
			return mapping.Format
		}
	}
	log.Fatalf("No valid Content-Type header present in server response. Does the resource URL support linked data? Header: %v", headerFixed)
	return rdf.NTriples
}

// acceptString -- Returns a string of all the available MIME types
func acceptString() string {
	str := ""
	for _, t := range rdfmimetypes.ContentTypes {
		str += t.Header
		str += " ,"
	}
	return str
}
