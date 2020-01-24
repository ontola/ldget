package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/kami-zh/go-capturer"

	srv "github.com/ontola/ldget/testing"
)

var subj = "https://app.argu.co/argu/u/joep"
var subjOut = "<https://app.argu.co/argu/u/joep>\n"
var pred = "http://schema.org/description"
var baseResource = "--resource=http://localhost:8080/joep.ttl"
var predOut = "<http://schema.org/description>\n"
var ntriple = "<https://app.argu.co/argu/u/joep> <http://schema.org/description> \"Liefhebber van discussiëren, ontwerpen en problemen oplossen. Een van de mede-oprichters van Argu.\" .\n"
var descriptionOut = "\"Liefhebber van discussiëren, ontwerpen en problemen oplossen. Een van de mede-oprichters van Argu.\"\n"
var predObjOut = "<http://schema.org/description> \"Liefhebber van discussiëren, ontwerpen en problemen oplossen. Een van de mede-oprichters van Argu.\"\n"

// Utrecht uses the metric system, right?
var extSubj = "http://dbpedia.org/resource/Utrecht"
var extPred = "http://dbpedia.org/property/metricFirst"
var extObjOut = "\"Yes\"@\n"

// Will probably return 'ldget'
var appname = os.Args[0:1][0]

var objectTests = []struct {
	in  []string
	out string
}{
	// Object ttl
	{[]string{appname, "objects", subj, pred, baseResource}, descriptionOut},
	// Object rdf, shortname
	{[]string{appname, "o", subj, pred, "--resource=http://localhost:8080/joep.rdf"}, descriptionOut},
	// Object nt
	{[]string{appname, "objects", subj, pred, "--resource=http://localhost:8080/joep.nt"}, descriptionOut},
	// Object prefix
	{[]string{appname, "o", subj, "schema:description", "--resource=http://localhost:8080/joep.rdf"}, descriptionOut},
	// TODO: add JSON-LD support
	// {[]string{appname, "objects", subj, pred, "--resource=http://localhost:8080/joep.jsonld"}, description},
	// Triples
	{[]string{appname, "triples", subj, pred, baseResource}, ntriple},
	{[]string{appname, "t", subj, pred, baseResource}, ntriple},
	// Subjects
	{[]string{appname, "subjects", subj, pred, baseResource}, subjOut},
	// Predicates
	{[]string{appname, "predicates", subj, pred, baseResource}, predOut},
	// PredicateObjects
	{[]string{appname, "po", subj, pred, baseResource}, predObjOut},
	// External resource
	{[]string{appname, "o", extSubj, extPred}, extObjOut},
}

func TestObjectParser(t *testing.T) {
	go srv.Testserver()

	// Run once to setup prefixes
	run([]string{appname, "o", extSubj, extPred})

	// Execute every single test string from objectTests
	for _, tt := range objectTests {
		fmt.Print(tt.in[0:])
		out := capturer.CaptureStdout(func() {
			run(tt.in)
		})
		if tt.out != out {
			t.Error(fmt.Sprintf("Expected: \n%vGot:\n%v", tt.out, out))
		} else {
			fmt.Print(" -- PASS\n")
		}
	}
}
