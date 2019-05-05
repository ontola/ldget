package main

import (
	"os"
	"testing"

	"github.com/kami-zh/go-capturer"
)

func TestRun(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "get")
	args = append(args, "--resource=https://app.argu.co/u/joep.nq")
	args = append(args, "--subject=https://app.argu.co/u/joep")
	args = append(args, "--predicate=http://schema.org/name")
	out := capturer.CaptureStderr(func() {
		run(args)
	})
	if out != "joe" {
		t.Error("Expected 'joe', got", out)
	}
}
