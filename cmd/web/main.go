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

	address := "localhost:4000"
	log.Printf("Starting Server on %s", address)
	err := http.ListenAndServe(address, mux)
	log.Fatal(err)
}
