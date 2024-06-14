package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Conditions")

	wg := &sync.WaitGroup{}

	scores := []int{0}

	go func(wg *sync.WaitGroup) {
		fmt.Println("One R")
		scores = append(scores, 1)
	}(wg)

	go func(wg *sync.WaitGroup) {
		fmt.Println("Two R")
		scores = append(scores, 2)
	}(wg)

	go func(wg *sync.WaitGroup) {
		fmt.Println("Three R")
		scores = append(scores, 3)
	}(wg)

}
