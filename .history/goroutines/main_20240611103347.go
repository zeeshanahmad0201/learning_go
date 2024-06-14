package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

var signals = []string{"test"}

func main() {
	websites := []string{
		"https://google.com",
		"https://fb.com",
		"https://linkedin.com",
	}

	for _, web := range websites {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	result, err := http.Get(endpoint)

	if err != nil {
		log.Printf("OOPs we got an error for %s saying\n%s\n", endpoint, err)
	}

	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
}
