package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/hello", myHandler)
  http.ListenAndServe(":8080", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-type", "text/plain")
  fmt.Fprintf(w, "Hello World")
}
