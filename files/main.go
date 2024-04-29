package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// create a file
	file, err := os.Create("./myFile.txt")
	checkError(err)

	// write to file
	length, err := io.WriteString(file, "This is my file")
	checkError(err)

	fmt.Println("length of the file is: ", length)

	// read file
	byteData, err := os.ReadFile("./myFile.txt")
	fmt.Println("This is in the file:\n", string(byteData))

	defer file.Close()

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	file, err := os.Create("./myfile.txt")
// 	checkNilErr(err)

// 	content := "This is my first file for golang created through golang code"

// 	length, err := io.WriteString(file, content)
// 	checkNilErr(err)

// 	fmt.Println("The length of the file is: ", length)

// 	defer file.Close()

// }

// func readFile(filename string) {
// 	byteData, err := os.ReadFile(filename)
// 	checkNilErr(err)

// 	fmt.Println("byteData is the following:\n", byteData)

// }

// func checkNilErr(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
