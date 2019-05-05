package main

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	args := os.Args[0:1]       // Name of the program.
	args = append(args, "get") // Append a flag
	run(args)
}
