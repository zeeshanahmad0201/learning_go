package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var sg *sync.WaitGroup

func main() {
	websites := []string{
		"https://google.com",
		"https://fb.com",
		"https://linkedin.com",
	}

	for _, web := range websites {
		getStatusCode(web)
		sg.Add(1)
	}
}

func getStatusCode(endpoint string) {
	defer sg.Done()
	result, err := http.Get(endpoint)

	if err != nil {
		log.Printf("OOPs we got an error for %s saying\n%s\n", endpoint, err)
	}

	fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
}
