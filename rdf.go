package main

// Triple -- A single triple
type Triple struct {
	subject   string
	object    string
	predicate string
}

// Finds all triples for a certain subject
func findBySubject(triples []*Triple, subject string) []*Triple {
	var hits []*Triple

	for _, t := range triples {
		if t.subject == subject {
			hits = append(hits, t)
		}
	}

	return hits
}

// Finds all triples with a certain predicate
func findByPredicate(triples []*Triple, predicate string) []*Triple {
	var hits []*Triple

	for _, t := range triples {
		if t.predicate == predicate {
			hits = append(hits, t)
		}
	}

	return hits
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
