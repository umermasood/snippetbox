package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// updated the signature of handler - now they are methods of *application
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// use the custom logger errorLog
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		// use the custom logger errorLog
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// updated the signature of handler - now they are methods of *application
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	_, _ = fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// updated the signature of handler - now they are methods of *application
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	_, _ = w.Write([]byte("Create a new snippet..."))
}
