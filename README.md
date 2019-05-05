# LDT: Linked Data Tool

A simple command line interface tool to get and manipulate RDF items.
Supports N-Quads.
Not yet ready for prime time, still in development.

## Run locally

`git clone https://bitbucket.org/joepio/argu-cli`
`go install`
`ld get --resource=https://argu.co/u/joep.ttl`

## Test

`go test`

## TODO

[] - Use a better parser. Currently, it only parses N-Quads, and it does so horribly.
