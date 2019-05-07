package main

import (
	"github.com/knakk/rdf"
	"net/http"
	// "strings"d
	"regexp"
)

// Negotiator -- Tries to fetch a resource using HTTP content negotiation
func Negotiator(url string) (*http.Response, rdf.Format, error) {
	// resp, err := http.Get(url)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", acceptString())
	resp, err := client.Do(req)
	respType := resp.Header.Get("Content-Type")
	format := findFormat(respType)

	return resp, format, err
}

type rdfFormatMapping struct {
	header string
	format rdf.Format
}

var acceptSelector, _ = regexp.Compile(`(.*);`)

func findFormat(header string) rdf.Format {
	matches := acceptSelector.FindStringSubmatch(header)
	headerFixed := header
	if len(matches) > 0 {
		headerFixed = matches[1]
	}
	for _, mapping := range contentTypes {
		if mapping.header == headerFixed {
			return mapping.format
		}
	}
	return rdf.NTriples
}

// acceptString -- Returns a string of all the available MIME types
func acceptString() string {
	str := ""
	for _, t := range contentTypes {
		str += t.header
		str += " ,"
	}
	return str
}

var contentTypes = []rdfFormatMapping{
	{
		header: "application/n-triples",
		format: rdf.NTriples,
	},
	{
		header: "application/rdf+xml",
		format: rdf.RDFXML,
	},
	{
		header: "application/x-turtle",
		format: rdf.Turtle,
	},
}
