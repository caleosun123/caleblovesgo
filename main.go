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
	mux.HandleFunc("/404", notFoundHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && r.URL.Path != "/message" {
			http.Redirect(w, r, "/404", http.StatusSeeOther)
		} else {
			handler(w, r)
		}
	})

	http.ListenAndServe(":8080", mux)
}
