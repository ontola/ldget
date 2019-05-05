# LDT: Linked Data Tool

A simple command line interface tool to get and manipulate RDF items.
Supports N-Quads.
Not yet ready for prime time, still in development.

## Run locally

`git clone https://bitbucket.org/joepio/argu-cli`
`go install`
`ld getObjects https://app.argu.co/argu/u/joep http://schema.org/description"`

## Mapping

You can specify a `mapping.ldmap` file for writing shorthands.
`ld o joep description"`

```
// in Mapping.ldmap
joep=https://app.argu.co/argu/u/joep
description=http://schema.org/description
```


## Test

`go test`

## TODO

[] - Use a better parser. Currently, it only parses N-Quads, and it does so horribly.
[] - Support JSON-LD @context objects, and map them for easy to use ORM.
