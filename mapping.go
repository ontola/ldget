package main

import (
	"log"
	"regexp"
)

// Map -- A user defined combination of prefixes
type Map struct {
	prefixes []prefix
}

type prefix struct {
	key string
	url string
}

var myFirstPrefix = prefix{"joep", "https://app.argu.co/argu/u/joep"}
var mySecondPrefix = prefix{"description", "http://schema.org/description"}

var prefixArray = []prefix{myFirstPrefix, mySecondPrefix}

func readMap() Map {
	return Map{prefixArray}
}

var myMap = readMap()

// Mapper -- converts a mappeed string to a URI
func Mapper(str string) string {
	matched, err := regexp.MatchString(`http.*`, str)
	if err != nil {
		log.Fatal(err)
	}
	output := str
	if matched {
		return output
	}
	for _, prefix := range myMap.prefixes {
		if prefix.key == str {
			output = prefix.url
			break
		}
	}
	return output
}
