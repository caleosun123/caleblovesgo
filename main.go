package main

import (
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, nil)
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from HTMX!"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tmpl := template.Must(template.ParseFiles("static/404.html"))
	tmpl.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/message", messageHandler)

	// Custom 404 handler
	mux.Handle("/404", http.HandlerFunc(notFoundHandler))
	mux.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", mux)
}
