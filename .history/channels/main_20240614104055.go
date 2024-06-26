package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myChan := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		myChan <- 5
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		fmt.Println(<-myChan)
		wg.Done()
	}(wg)

	wg.Wait()
}
