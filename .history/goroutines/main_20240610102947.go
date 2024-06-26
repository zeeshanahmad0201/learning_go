package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	websites := []string{
		"https://gooogle.com",
		"https://fb.com",
		"https://linkedin.com"
	}
}

func getStatusCode(endpoint string) {
	result, err := http.Get(endpoint)

	if err != nil {
		log.Printf("OOPs we got an error for %s saying\n%s", endpoint, err)
	}

	fmt.Printf("%d status code for %s", result.StatusCode, endpoint)
}
