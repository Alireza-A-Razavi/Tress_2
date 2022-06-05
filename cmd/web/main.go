package main

import (
  "log"
  "net/http"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet", showSnippet)
  mux.HandleFunc("/snippet/create", createSnippet)

  // static files
  fileServer := http.FileServer(http.Dir("./ui/static/"))
  // NOTE: // you don't want to use HandleFunc here!
  mux.Handle("/static", http.SrtipPrefix("/static", fileServer))
  
  log.Println("Starting server on :4000")
  error := http.ListenAndServer(":4000", mux)
  log.Fatal(error)
}
