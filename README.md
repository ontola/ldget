# LDT: Linked Data Tool

A simple command line interface tool to get and manipulate RDF items.
Supports N-Quads.
Not yet ready for prime time, still in development.

## Run locally

`git clone https://bitbucket.org/joepio/argu-cli`
`go install`
`ldget getObjects https://app.argu.co/argu/u/joep http://schema.org/description"`

## Mapping

You can specify a `mapping.ldget` file for writing shorthands.
`ldget getObjects joep description"`

```
// in mapping.ldmap
joep=https://app.argu.co/argu/u/joep
description=http://schema.org/description
```

## Test

`go test`

## TODO

[] - Use a better parser. Currently, it only parses N-Quads, and it does so horribly.
[] - Support JSON-LD @context objects, and map them for easy to use ORM.
[] - Use content negotiation / accept headers
[] - Traverse relationships, fetch content across websites.
