package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/kami-zh/go-capturer"
)

func TestRun(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "get")
	args = append(args, "--resource=https://app.argu.co/u/joep.nq")
	args = append(args, "--subject=https://app.argu.co/argu/u/joep")
	args = append(args, "--predicate=http://schema.org/description")
	out := capturer.CaptureStdout(func() {
		run(args)
	})
	value := "\"Liefhebber van discussiëren, ontwerpen en problemen oplossen. Een van de mede-oprichters van Argu.\"\n"
	if out != value {
		t.Error(fmt.Sprintf("Expected: \n%vGot:\n%v", value, out))
	}
}
