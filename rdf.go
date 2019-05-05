package main

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
