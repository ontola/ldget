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

type prefix struct {
	key string
	url string
}

// Mapper -- converts a mapped prefix to a URI
func Mapper(str string) string {
	httpCheck, err := regexp.MatchString(`http.*`, str)
	if err != nil {
		log.Fatal(err)
	}
	output := str
	// If the input starts with http, don't look up the mapping
	if httpCheck {
		return output
	}

	// Check for colon prefix syntax, e.g. `schema:description`
	matches := colonCheck.FindStringSubmatch(str)
	if len(matches) > 2 {
		output = fmt.Sprintf("%v%v", getPrefix(matches[1]), matches[2])
	} else {
		// Directly use the prefix
		output = getPrefix(str)
	}

	return output
}

// Regex for the user's ~/.ldget/prefixes file
var selector, _ = regexp.Compile(`(.*)=(.*)`)

func readMap(filePath string) []prefix {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var prefixes []prefix
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Lines that start with # are comments
		if strings.HasPrefix(line, "#") {
			continue
		}
		matches := selector.FindStringSubmatch(line)
		if len(matches) < 2 {
			log.Fatal("Something is wrong with your prefixes file.")
		}
		var p prefix
		p.key = matches[1]
		p.url = matches[2]
		prefixes = append(prefixes, p)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return prefixes
}

func getAllMaps() []prefix {
	var allPrefixes []prefix

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	userMappingLocation := fmt.Sprintf("%v/.ldget/prefixes", usr.HomeDir)

	allPrefixes = append(allPrefixes, readMap(userMappingLocation)...)
	return allPrefixes
}

var colonCheck, _ = regexp.Compile(`(.*):(.*)`)

// Returns URL for some prefix
func getPrefix(key string) string {
	output := key
	for _, prefix := range getAllMaps() {
		if prefix.key == key {
			output = prefix.url
			break
		}
	}
	return output
}
