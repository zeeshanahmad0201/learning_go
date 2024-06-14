package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myChan := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(wg *sync.WaitGroup, ch chan int) {
		ch <- 5
		ch <- 6
		wg.Done()
	}(wg, myChan)

	go func(wg *sync.WaitGroup, ch chan int) {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		wg.Done()
	}(wg, myChan)

	wg.Wait()
}
