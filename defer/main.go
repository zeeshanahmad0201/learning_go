package main

import "fmt"

func main() {
	defer fmt.Println("Here is deferred")
	defer myDefer()
	fmt.Println("Hello World!")
}

// hello world!, 4, 3, 2, 1, 0, Here is deferred
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
