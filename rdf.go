package main

// Triple -- A single triple
type Triple struct {
	subject   string
	object    string
	predicate string
}

// Filter triples by subject, predicate, object
func filterTriples(triples []*Triple, subject string, predicate string, object string) []*Triple {
	var hits []*Triple

	for _, t := range triples {
		var doTheyMatch = (((subject == "") || (t.subject == subject)) &&
			((object == "") || (t.object == object)) &&
			((predicate == "") || (t.predicate == predicate)))
		if doTheyMatch {
			hits = append(hits, t)
		}
	}

	return hits
}
