package main

import (
	"fmt"
	"http"
	"log"
)

func main() {
	// Create a file server handler
	fileServer := http.FileServer(http.Dir("./static"))

	// Handle all requests to the web root with the file server
	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nill); err != nil {
		log.Fatal(err)
	}

}
