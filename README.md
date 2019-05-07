# ldget
[![Build Status](https://travis-ci.org/ontola/active_response.svg?branch=master)](https://travis-ci.org/ontola/active_response) ![GitHub](https://img.shields.io/github/license/ontola/ldget.svg)

A simple command line interface tool to get RDF items using HTTP GET requests.

## Installation

- On MacOS using [homebrew](https://brew.sh/): `$ brew tap ontola/ldget https://github.com/ontola/ldget.git && brew install ontola/ldget`

## Usage

- `$ ldget object ${subjectIRI} ${predicateIRI}` => returns the values of the objects that match
- `$ ldget help` => help file
- `$ ldget prefixes` => shows your configured prefixes

## Why should you use this?

- You need RDF data as Stdout.
- You want to write bash scripts that use linked data.
- You need to check triple values from inside your terminal.

## Prefixes

URLs are awesome, but they are cumbersome to remember and type.
You can specify a set of prefixes in `~/.ldget/prefixes` for mapping URLS to shorthands.

```
schema=http://schema.org/
joep=https://argu.co/argu/u/joep
```

`$ ldget objects joep schema:description`

## Install

- `$ git clone git@github.com:ontola/ldget.git`
- `$ go install`

## Test

- `$ go test`
