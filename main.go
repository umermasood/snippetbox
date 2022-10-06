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
	// throw 405 (method not allowed) if the request method is not POST
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("Method Not Allowed"))
		return
	}
	_, _ = w.Write([]byte("Create a new snippet..."))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/snippet/create", snippetCreate)
	http.HandleFunc("/snippet/view", snippetView)

	log.Println("Starting the web server on port :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
