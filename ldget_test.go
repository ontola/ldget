package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/kami-zh/go-capturer"
)

func setArgs() []string {
	args := os.Args[0:1]
	args = append(args, "objects")
	args = append(args, "https://app.argu.co/argu/u/joep")
	args = append(args, "http://schema.org/description")
	return args
}

var description = "Liefhebber van discussiëren, ontwerpen en problemen oplossen. Een van de mede-oprichters van Argu.\n"

var objectTests = []struct {
	in  []string
	out string
}{
	{setArgs(), description},
}

func TestObjectParser(t *testing.T) {
	for _, tt := range objectTests {
		out := capturer.CaptureStdout(func() {
			run(tt.in)
		})
		if tt.out != out {
			t.Error(fmt.Sprintf("Expected: \n%vGot:\n%v", tt.out, out))
		}
	}
}
