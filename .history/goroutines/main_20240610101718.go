package main

import "fmt"

func main() {

}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}