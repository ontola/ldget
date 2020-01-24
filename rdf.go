package main

import (
	"io"
	"log"

	"github.com/knakk/rdf"
)

// Filter triples by subject, predicate, object
func filterTriples(triples []rdf.Triple, subject string, predicate string, object string) []rdf.Triple {
	var hits []rdf.Triple

	for _, t := range triples {
		var doTheyMatch = (((subject == "") || (t.Subj.String() == subject)) &&
			((object == "") || (t.Obj.String() == object)) &&
			((predicate == "") || (t.Pred.String() == predicate)))
		if doTheyMatch {
			hits = append(hits, t)
		}
	}

	return hits
}

// Parse -- Reads a stream of NQuads and returns an array of Triples
func Parse(body io.Reader, format rdf.Format) ([]rdf.Triple, error) {
	decoder := rdf.NewTripleDecoder(body, format)
	var triples []rdf.Triple
	for triple, err := decoder.Decode(); err != io.EOF; triple, err = decoder.Decode() {
		triples = append(triples, triple)
	}
	return triples, nil
}

// Fetches triples for a set of arguments, returns the filtered triples
func getTriples(a Args) []rdf.Triple {
	resp, format, err := Negotiator(a.resourceURL, a)
	if err != nil {
		log.Fatal(err)
	}
	// Parse the response, convert to triples
	allTriples, err := Parse(resp.Body, format)
	if err != nil {
		log.Fatal(err)
	}
	hits := filterTriples(allTriples, a.subject, a.predicate, a.object)
	if len(hits) == 0 {
		log.Fatal("No triple found that matches your query")
	}
	return hits
}
