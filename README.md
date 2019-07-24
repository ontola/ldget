# ldget
[![Build Status](https://travis-ci.org/ontola/active_response.svg?branch=master)](https://travis-ci.org/ontola/active_response) ![GitHub](https://img.shields.io/github/license/ontola/ldget.svg)

A simple command line interface tool to get RDF items using HTTP GET requests.

[![asciicast](demo.gif)](https://asciinema.org/a/BOc4JjKkTbMOwA2IvnJ7V3uhp)

## When should you use this?

- You need RDF data as Stdout.
- You want to write bash scripts that use linked data.
- You need to check triple values from inside your terminal.

## Installation

- On MacOS using [homebrew](https://brew.sh/): `$ brew tap ontola/ldget https://github.com/ontola/ldget.git && brew install ontola/ldget/ldget`
- Or install the binaries from the [releases page](https://github.com/ontola/ldget/releases)
- Or build it yourself. Clone this repo and run `go get ./... && go install` to install the dependencies and the binaries.

## Usage

- `$ ldget triples ?s ?p ?o` => fetches the subject (?s) URL, returns all triples that match
- `$ ldget predicates ?s ?p ?o` => fetches the subject (?s) URL, returns the predicates that match
- `$ ldget objects ?s ?p ?o` => fetches the subject (?s) URL, returns the objects that match
- `$ ldget help` => help file
- `$ ldget prefixes` => shows your configured prefixes

`?s ?p ?o` stands for `%{subject URL} %{predicate URL} %{object value}`. This is similar to [Triple Pattern Fragments](http://www.hydra-cg.com/spec/latest/triple-pattern-fragments/#bib-hydra-ldf).

Use the `*` character as a wildcard. For example, if you want to get all triples for subject `http://example.com/X` with an object value of `"Value"`, use:

`$ ldget t http://example.com/X * "Value"`

## Prefixes

URLs are awesome, but they are cumbersome to remember and type.
You can specify a set of prefixes in `~/.ldget/prefixes` for mapping URLs to shorthands.

```
schema=http://schema.org/
joep=https://argu.co/argu/u/joep
```

`$ ldget o joep schema:description`

## Install

- `$ git clone git@github.com:ontola/ldget.git && cd ldget` Clone repo, go to ldget folder
- `$ go get ./...` Install all dependencies
- `$ go install` Install binary to path

## Test

- `$ go test`

## Credits

Written by Joep Meindertsma.

Most of the hard work is done by the guys at Knakk, who wrote [this awesome RDF library for Go](https://github.com/knakk/rdf).
