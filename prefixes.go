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

var colonCheck, _ = regexp.Compile(`(.*):(.*)`)

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
