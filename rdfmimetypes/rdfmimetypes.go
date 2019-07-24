package rdfmimetypes

import "github.com/knakk/rdf"

// RdfFormatMapping - An RDF format
type RdfFormatMapping struct {
	// The mime header, e.g. `application/rdf+xml`
	Header string
	// Includes the ., e.g. `.rdf`
	Extension string
	// The Knakk.RDF format
	Format rdf.Format
}

// ContentTypes -- List of supported RDF types
var ContentTypes = []RdfFormatMapping{
	{
		Header:    "application/n-triples",
		Extension: ".nt",
		Format:    rdf.NTriples,
	},
	{
		Header:    "application/rdf+xml",
		Extension: ".rdf",
		Format:    rdf.RDFXML,
	},
	{
		Header:    "text/turtle",
		Extension: ".ttl",
		Format:    rdf.Turtle,
	},
}
