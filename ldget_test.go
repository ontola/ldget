package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/kami-zh/go-capturer"
)

var description = "\"Liefhebber van discussiëren, ontwerpen en problemen oplossen. Een van de mede-oprichters van Argu.\"\n"

func TestGetObject(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "objects")
	args = append(args, "--resource=https://app.argu.co/u/joep.nq")
	args = append(args, "--subject=https://app.argu.co/argu/u/joep")
	args = append(args, "--predicate=http://schema.org/description")
	out := capturer.CaptureStdout(func() {
		run(args)
	})
	if out != description {
		t.Error(fmt.Sprintf("Expected: \n%vGot:\n%v", description, out))
	}
}

func TestGetObjectArgs(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "objects")
	args = append(args, "https://app.argu.co/argu/u/joep")
	args = append(args, "http://schema.org/description")
	out := capturer.CaptureStdout(func() {
		run(args)
	})
	if out != description {
		t.Error(fmt.Sprintf("Expected: \n%vGot:\n%v", description, out))
	}
}
