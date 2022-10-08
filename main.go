package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, _ = w.Write([]byte("Hello from snippet box"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	_, _ = fmt.Fprintf(w, "Display a specific snippet by ID %d...", id)
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

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/snippet/create", snippetCreate)
	http.HandleFunc("/snippet/view", snippetView)

	log.Println("Starting the web server on port :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
