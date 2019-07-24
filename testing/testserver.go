package testserver

import (
	"bufio"
	"net/http"
	"os"
	"strings"

	rdfmimetypes "github.com/ontola/ldget/rdfmimetypes"
)

type myHandler struct {
	http.Handler
}

// Check the extension of the file, return a contentType.
func contentType(path string) (contentType string) {
	for _, mapping := range rdfmimetypes.ContentTypes {
		if strings.HasSuffix(path, mapping.Extension) {
			contentType = mapping.Header
		}
	}
	return contentType
}

// We create a custom handler
func (handler *myHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "./testing/" + req.URL.Path
	// Open a file, which does not render it but keep it ready for read
	f, err := os.Open(path)

	// if a file exists, check its content type, or return 404
	if err == nil {
		// read the content to buffer in order to save memory
		bufferedReader := bufio.NewReader(f)
		// check content type of the file according to its suffix
		w.Header().Add("Content-Type", contentType(path))
		// write the file content to the response
		bufferedReader.WriteTo(w)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}

// Testserver -- HTTP server that handles some RDF requests
func Testserver() {
	// use the custom handler
	http.Handle("/", new(myHandler))
	http.ListenAndServe(":8080", nil)
}
