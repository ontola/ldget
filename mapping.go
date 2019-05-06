package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"regexp"
)

type prefix struct {
	key string
	url string
}

var selector, _ = regexp.Compile(`(.*)=(.*)`)

func readMap(filePath string) []prefix {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var prefixes []prefix
	scanner := bufio.NewScanner(file) // f is the *os.File
	for scanner.Scan() {
		line := scanner.Text()
		matches := selector.FindStringSubmatch(line)
		var p prefix
		if len(matches) < 2 {
			log.Fatal("Something is wrong with your mapping file.")
		}
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
	allPrefixes = append(allPrefixes, readMap("defaultMapping")...)

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	userMappingLocation := fmt.Sprintf("%v/.ldget/mapping", usr.HomeDir)

	allPrefixes = append(allPrefixes, readMap(userMappingLocation)...)
	return allPrefixes
}

var colonCheck, _ = regexp.Compile(`(.*):(.*)`)

// Mapper -- converts a mapped string to a URI
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
