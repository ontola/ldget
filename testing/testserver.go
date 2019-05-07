package testserver

import (
	"net/http"
)

// Testserver -- HTTP server that handles some RDF requests
func Testserver() {
	http.HandleFunc("/ping", ping)
	http.Handle("/", http.FileServer(http.Dir("./testing/")))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
