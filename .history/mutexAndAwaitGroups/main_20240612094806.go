package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Conditions")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	scores := []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("One R")
		scores = append(scores, 1)
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Two R")
		scores = append(scores, 2)
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Three R")
		scores = append(scores, 3)
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(scores)
}
