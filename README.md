# LDT: Linked Data Tool

A simple command line interface tool to get and manipulate RDF items.
Only supports N-Quads.
Not yet ready for prime time, still in development.

## Run locally

`git clone https://bitbucket.org/joepio/argu-cli`
`go install`
`ldget objects https://app.argu.co/argu/u/joep http://schema.org/description"`

## Mapping

You can specify an `~/.ldget/mapping` file for writing shorthands / @prefixes.

```
// in ~/.ldget/mapping
joep=https://app.argu.co/argu/u/joep
description=http://schema.org/description
```

If you have a mapping, you can use shorthand prefixes: `ldget objects joep description"`.

## Test

`go test`

## TODO

[] - Use a better parser. Currently, it only parses N-Quads, and it does so horribly.
[] - Support (external) JSON-LD @context objects, and map them for easy to use ORM.
[] - Prefix colon syntax (e.g. schema:description)
[] - Use content negotiation / accept headers
[] - Traverse relationships, fetch content across websites.
