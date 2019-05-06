package main

import (
	"io"

	"github.com/knakk/rdf"
)

// Parse -- Reads a stream of NQuads and returns an array of Triples
func Parse(body io.Reader) ([]rdf.Triple, error) {
	decoder := rdf.NewTripleDecoder(body, rdf.NTriples)
	var triples []rdf.Triple
	for triple, err := decoder.Decode(); err != io.EOF; triple, err = decoder.Decode() {
		triples = append(triples, triple)
	}
	return triples, nil
}
