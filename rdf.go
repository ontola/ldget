package main

import (
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
