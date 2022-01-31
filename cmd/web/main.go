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

<<<<<<< HEAD
	// Create a file server which serves files out of the "./ui/static" directory
	// Note that the path given to the http.Dir function is relative to the project directory root.
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Use the mux.Handle() function to register the file server as the handler for all URL paths that start with "/static". For matching paths, we strip the "/static" prefix before the request reaches the file server
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on localhost:4000")
	err := http.ListenAndServe(":4000", mux)
=======
	address := "localhost:4000"
	log.Printf("Starting Server on %s", address)
	err := http.ListenAndServe(address, mux)
>>>>>>> 4755cf1ab950a92e870efea7b3f34b34e03e5a09
	log.Fatal(err)
}
