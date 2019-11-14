package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/knakk/rdf"

	// "strings"d
	"log"
	"regexp"

	rdfmimetypes "github.com/ontola/ldget/rdfmimetypes"
)

// Negotiator -- Tries to fetch a resource using HTTP content negotiation
func Negotiator(url string, args args) (*http.Response, rdf.Format, error) {
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
	verbose := args.verbose
	if verbose == true {
		// Debugging puposes
		fmt.Println(formatRequest(req))
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

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
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
		str += ", "
	}
	return str
}
