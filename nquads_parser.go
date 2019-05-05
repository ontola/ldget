package main

import (
	"bufio"
	"io"
	"regexp"
)

var fishRemover, _ = regexp.Compile(`<(.+?)>`)
var nQuadsRegex, _ = regexp.Compile(`(<.+?>) (<.+?>) (.+?) \.`)

// NewTriple -- Creates a triple from three strings
func NewTriple(subject string, predicate string, object string) *Triple {
	triple := new(Triple)
	triple.object = object
	triple.predicate = removeFishHooks(predicate)
	triple.subject = removeFishHooks(subject)
	return triple
}

func removeFishHooks(iri string) string {
	matches := fishRemover.FindStringSubmatch(iri)
	return matches[1]
}

// Parse -- Reads a stream of NQuads and returns an array of Triples
func Parse(body io.Reader) (triples []*Triple, err error) {
	scanner := bufio.NewScanner(bufio.NewReader(body))
	// Iterates over every single triple
	var parsedTriples []*Triple
	for scanner.Scan() {
		triple := scanner.Text()
		matches := nQuadsRegex.FindStringSubmatch(triple)
		createdTriple := NewTriple(matches[1], matches[2], matches[3])
		parsedTriples = append(parsedTriples, createdTriple)
	}
	return parsedTriples, nil
}
