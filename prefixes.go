package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"regexp"
	"strings"
)

// Prefix - A single combination of a key and URL
type Prefix struct {
	// The short version, e.g. "schema"
	key string
	// The long version, e.g. "https://schema.org/"
	url string
}

// prefixToURL - Converts a prefix string to a full URI.
// Returns the input string if no prefix is found.
func prefixToURL(str string, prefixes []Prefix) string {
	httpCheck, err := regexp.MatchString(`http.*`, str)
	if err != nil {
		log.Fatal(err)
	}
	// By default, return the input string
	output := str
	// If the input starts with http, don't look up the mapping
	if httpCheck {
		return output
	}

	// Check for colon prefix syntax, e.g. `schema:description`
	matches := colonCheck.FindStringSubmatch(str)
	if len(matches) > 2 {
		output = fmt.Sprintf("%v%v", getPrefix(matches[1], prefixes), matches[2])
	} else {
		// Directly use the prefix
		output = getPrefix(str, prefixes)
	}

	return output
}

// Regex for the user's ~/.ldget/prefixes file
var selector, _ = regexp.Compile(`(.*)=(.*)`)

// Parses the prefixes file
func readPrefixesFile(filePath string) []Prefix {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var prefixes []Prefix
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Lines that start with # are comments
		if strings.HasPrefix(line, "#") {
			continue
		}
		// Ignore empty lines
		if line == "" {
			continue
		}
		matches := selector.FindStringSubmatch(line)
		if len(matches) < 2 {
			log.Fatal("Something is wrong with your prefixes file.")
		}
		var p Prefix
		p.key = matches[1]
		p.url = matches[2]
		prefixes = append(prefixes, p)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return prefixes
}

func getAllPrefixes() []Prefix {
	var allPrefixes []Prefix

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	userMappingLocation := fmt.Sprintf("%v/.ldget/prefixes", usr.HomeDir)

	allPrefixes = append(allPrefixes, readPrefixesFile(userMappingLocation)...)
	return allPrefixes
}

func removeBrackets(url string) string {
	noSuffix := strings.TrimSuffix(url, ">")
	output := strings.TrimPrefix(noSuffix, "<")
	return output
}

// Converts a URL to a prefix
// e.g. <https://schema.org/hello> => schema:hello
func tryURLToPrefix(url string, args Args) string {
	cleanURL := removeBrackets(url)
	output := cleanURL

	for _, prefix := range args.prefixes {
		if strings.HasPrefix(cleanURL, prefix.url) {
			prefixWithColon := fmt.Sprintf("%v:", prefix.key)
			newString := strings.Replace(cleanURL, prefix.url, prefixWithColon, 1)
			output = newString
			break
		}
	}
	return output
}

var colonCheck, _ = regexp.Compile(`(.*):(.*)`)

// Returns URL for some prefix
func getPrefix(key string, prefixes []Prefix) string {
	output := key
	for _, prefix := range prefixes {
		if prefix.key == key {
			output = prefix.url
			break
		}
	}
	return output
}
