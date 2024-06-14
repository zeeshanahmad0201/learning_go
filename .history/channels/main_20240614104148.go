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
	go func(wg *sync.WaitGroup, ch chan int) {
		ch <- 5
		wg.Done()
	}(wg, myChan)

	go func(wg *sync.WaitGroup, ch chan int) {
		fmt.Println(<-myChan)
		wg.Done()
	}(wg, myChan)

	wg.Wait()
}
