package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Hello from snippet box"))
}

// Add a snippetView handler function
func snippetView(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Display a specific snippet..."))
}

// Add a snippetCreate handler function
func snippetCreate(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	// registering two new handler functions and corresponding URL patterns
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/snippet/view", snippetView)

	log.Println("Starting the web server on port :4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}
