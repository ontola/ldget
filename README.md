# ldget
[![Build Status](https://travis-ci.org/ontola/active_response.svg?branch=master)](https://travis-ci.org/ontola/active_response)

A simple command line interface tool to get RDF items using HTTP GET requests.

Not yet ready for prime time, still in development.

## When should you use this?

- You need RDF data as Stdout.
- You want to write bash scripts that use linked data.
- You need to check triple values quite often from inside your terminal.

## Run locally

- `git clone https://bitbucket.org/joepio/argu-cli`
- `go install`
- `ldget objects https://app.argu.co/argu/u/joep http://schema.org/description"`

## Prefixes

URLs are awesome, but they are cumbersome to remember and type.
You can specify a set of prefixes in `~/.ldget/prefixes` for mapping URLS to shorthands.

```
schema=http://schema.org/
```

`ldget objects https://argu.co/argu/u/joep schema:description`.

## Test

`go test`

## TODO

- [ ] Support (external) JSON-LD @context objects, and map them for easy to use ORM.
- [ ] Use locally hosted documents for testing
- [ ] Traverse relationships, fetch content across websites.
- [x] Use content negotiation / accept headers
- [x] Prefix colon syntax (e.g. schema:description)
- [x] Use a better parser. Currently, it only parses N-Quads, and it does so horribly.
- [x] Use table tests
