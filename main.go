package main

import (
	"fmt"
	"html"
	"net/http"
)

func serveGat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/gat", serveGat)

	fmt.Printf("Running server on port 8080\n")
	http.ListenAndServe(":8080", nil)
}
