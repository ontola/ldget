# ldget
[![Build Status](https://travis-ci.org/ontola/active_response.svg?branch=master)](https://travis-ci.org/ontola/active_response) ![GitHub](https://img.shields.io/github/license/ontola/ldget.svg)

A simple command line interface tool to get RDF items using HTTP GET requests.

## Usage

- `ldget object ${subjectIRI} ${predicateIRI}` => returns the values of the objects that match

## Why should you use this?

- You need RDF data as Stdout.
- You want to write bash scripts that use linked data.
- You need to check triple values from inside your terminal.

## Prefixes

URLs are awesome, but they are cumbersome to remember and type.
You can specify a set of prefixes in `~/.ldget/prefixes` for mapping URLS to shorthands.

```
schema=http://schema.org/
```

`ldget objects https://argu.co/argu/u/joep schema:description`.

## Install

- `git clone https://bitbucket.org/joepio/argu-cli`
- `go install`

## Test

`go test`
