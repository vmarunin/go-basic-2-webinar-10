package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		log.Println("Requested path:", req.URL.Path)
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/", helloHandler)
	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
