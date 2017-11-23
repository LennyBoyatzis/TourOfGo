package main

import (
	"fmt"
	"net/http"
)

// * Reading through http.Request
func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, go is awesome")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About this yall</h1>")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil)
}
