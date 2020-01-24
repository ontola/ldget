package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/user"
	"regexp"
	"strings"
)

var defaultPrefixes = []Prefix{
	Prefix{
		key: "schema",
		url: "http://schema.org/",
	},
	Prefix{
		key: "rdf",
		url: "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
	},
	Prefix{
		key: "owl",
		url: "http://www.w3.org/2002/07/owl#",
	},
	Prefix{
		key: "ldp",
		url: "http://www.w3.org/ns/ldp#",
	},
	Prefix{
		key: "rdfa",
		url: "http://www.w3.org/ns/rdfa#",
	},
	Prefix{
		key: "rdfs",
		url: "http://www.w3.org/2000/01/rdf-schema#",
	},
	Prefix{
		key: "skos",
		url: "http://www.w3.org/2004/02/skos/core#",
	},
	Prefix{
		key: "vcard",
		url: "http://www.w3.org/2006/vcard/ns#",
	},
	Prefix{
		key: "xsd",
		url: "http://www.w3.org/2001/XMLSchema#",
	},
	Prefix{
		key: "dbp",
		url: "http://dbpedia.org/",
	},
	Prefix{
		key: "dbpr",
		url: "http://dbpedia.org/resource/",
	},
	Prefix{
		key: "dbpd",
		url: "http://dbpedia.org/resource/",
	},
	Prefix{
		key: "dbpp",
		url: "http://dbpedia.org/property/",
	},
	Prefix{
		key: "ncal",
		url: "http://www.semanticdesktop.org/ontologies/2007/04/02/ncal#",
	},
	Prefix{
		key: "org",
		url: "http://www.w3.org/ns/org#",
	},
}

// Prefix - A single combination of a key and URL
type Prefix struct {
	// The short version, e.g. "schema"
	key string
	// The long version, e.g. "https://schema.org/"
	url string
}

func getAllPrefixes(path string) []Prefix {
	var allPrefixes []Prefix
	allPrefixes = append(allPrefixes, readPrefixesFile(path)...)
	return allPrefixes
}

// Parses the prefixes file
func readPrefixesFile(filePath string) []Prefix {
	file, err := os.Open(filePath)
	if err != nil {
		if strings.Contains(err.Error(), "no such file") {
			setDefaultPrefixes(filePath)
			// Open the newly created file
			file2, err2 := os.Open(filePath)
			if err2 != nil {
				log.Fatal(err2)
			}
			file = file2
		} else {
			log.Fatal(err)
		}
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

// Creates file with default prefixes at default location
func setDefaultPrefixes(filePath string) {
	log.Printf("No prefixes file found, creating one at %v", filePath)
	writePrefixes(filePath, defaultPrefixes)
}

// Writes prefixes to file
func writePrefixes(filePath string, prefixes []Prefix) {
	newfile, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.Buffer{}
	for _, element := range prefixes {
		buf.WriteString(element.key)
		buf.WriteString("=")
		buf.WriteString(element.url)
		buf.WriteString("\n")
	}
	result := buf.String()
	newfile.WriteString(result)
}

// GetPrefixPath - Retruns the path where .ldget/prefixes should live
func GetPrefixPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v/.ldget/prefixes", usr.HomeDir)
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
