package main

import "fmt"

func main() {
	fmt.Println("Race Conditions")

	scores := []int{0}

	go func() {
		fmt.Println("One R")
		scores = append(scores, 1)
	}()

	go func() {
		fmt.Println("Two R")
		scores = append(scores, 2)
	}()

	go func() {
		fmt.Println("Three R")
		scores = append(scores, 3)
	}()

}
