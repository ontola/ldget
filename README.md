# LDT: Linked Data Tool

A simple command line interface tool to get and manipulate RDF items.
Only supports N-Quads.
Not yet ready for prime time, still in development.

## Run locally

`git clone https://bitbucket.org/joepio/argu-cli`
`go install`
`ldget objects https://app.argu.co/argu/u/joep http://schema.org/description"`

## Prefixes

URLs are awesome, but they are cumbersome to remember and type.
You can specify a set of prefixes in `~/.ldget/prefixes` for mapping URLS to shorthands.

```
// in ~/.ldget/prefixes
joep=https://app.argu.co/argu/u/joep
schema=http://schema.org/
```

If you have a prefixes, you can use shorthand prefixes: `ldget objects joep schema:description"`.

## Test

`go test`

## TODO

[] - Use a better parser. Currently, it only parses N-Quads, and it does so horribly.
[] - Support (external) JSON-LD @context objects, and map them for easy to use ORM.
[x] - Prefix colon syntax (e.g. schema:description)
[] - Use content negotiation / accept headers
[] - Traverse relationships, fetch content across websites.
