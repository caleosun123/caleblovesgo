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

func main() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/message", messageHandler)
  http.ListenAndServe(":8080", nil)
}
