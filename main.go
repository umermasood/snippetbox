package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, _ = w.Write([]byte("Hello from snippet box"))
}

func snippetView(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/snippet/view", snippetView)

	log.Println("Starting the web server on port :4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}
