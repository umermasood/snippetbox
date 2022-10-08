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

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"response":"Create a new snippet..."}`))
}

func manipulateHandler(w http.ResponseWriter, _ *http.Request) {
	// Set a new cache-control header. If an existing "Cache-Control" header exists
	// it will be overwritten.
	w.Header().Set("cache-control", "public, max-age=31536000")
	// In contrast, the Add() method appends a new "Cache-Control" header and can
	// be called multiple times.
	w.Header().Add("Cache-Control", "public")
	w.Header().Add("Cache-Control", "max-age=31536000")
	w.Header()["foo-header"] = []string{"bar-val1"}
	// Delete all values for the "Cache-Control" header.
	//w.Header().Del("Cache-Control")
	// Retrieve the first value for the "Cache-Control" header.
	log.Println(w.Header().Get("Cache-Control"))
	// Retrieve a slice of all values for the "Cache-Control" header.
	log.Println(w.Header().Values("Cache-Control"))

	_, _ = w.Write([]byte("write to manipulate header"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/snippet/create", snippetCreate)
	http.HandleFunc("/snippet/view", snippetView)
	http.HandleFunc("/snippet/manipulate/header", manipulateHandler)

	log.Println("Starting the web server on port :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
