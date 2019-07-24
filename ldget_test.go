package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/kami-zh/go-capturer"

	srv "github.com/ontola/ldget/testing"
)

var description = "Liefhebber van discussiëren, ontwerpen en problemen oplossen. Een van de mede-oprichters van Argu.\n"
var ntriple = "<https://app.argu.co/argu/u/joep> <http://schema.org/description> \"Liefhebber van discussiëren, ontwerpen en problemen oplossen. Een van de mede-oprichters van Argu.\" .\n"

// Will probably return ldget
var appname = os.Args[0:1][0]

var objectTests = []struct {
	in  []string
	out string
}{
	{[]string{appname, "objects", "https://app.argu.co/argu/u/joep", "http://schema.org/description", "--resource=http://localhost:8080/joep.ttl"}, description},
	{[]string{appname, "objects", "https://app.argu.co/argu/u/joep", "http://schema.org/description", "--resource=http://localhost:8080/joep.rdf"}, description},
	{[]string{appname, "objects", "https://app.argu.co/argu/u/joep", "http://schema.org/description", "--resource=http://localhost:8080/joep.nt"}, description},
	// {[]string{appname, "objects", "https://app.argu.co/argu/u/joep", "http://schema.org/description", "--resource=http://localhost:8080/joep.jsonld"}, description},
	{[]string{appname, "triples", "https://app.argu.co/argu/u/joep", "http://schema.org/description", "--resource=http://localhost:8080/joep.ttl"}, ntriple},
	{[]string{appname, "predicates", "https://app.argu.co/argu/u/joep", "http://schema.org/description", "--resource=http://localhost:8080/joep.ttl"}, "http://schema.org/description\n"},
}

func TestObjectParser(t *testing.T) {
	go srv.Testserver()

	// Execute every single test string from objectTests
	for _, tt := range objectTests {
		fmt.Print(tt.in[1:])
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
