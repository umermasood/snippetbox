package main

import (
	"log"
	"net/http"
)

// Three Essentials of a Web Application
// 1. handler
// 2. router (called ServeMux in Go Terminology)
// 3. web server (go has a built-in web server, so we don't need third party web servers like Apache or NGINX)

// 1. handler
func home(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Hello from snippet box"))
}

func main() {
	// 2. router / serveMux
	mux := http.NewServeMux()

	// registering home function as handler for "/" URL pattern
	mux.HandleFunc("/", home)

	log.Println("Starting the web server on port :4000")

	// 3. web server
	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatal(err)
	}
}
