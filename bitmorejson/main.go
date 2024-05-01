package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Title    string   `json:"title"`
	Price    int      `json:"price"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to a bitmorejso")

	EncodeJson()

}

func EncodeJson() {

	courses := []course{
		{"flutter development", 199, "abc", []string{"flutter", "dev"}},
		{"python development", 100, "123", nil},
	}

	finalJson, err := json.MarshalIndent(courses, "", "\t")
	checkErr(err)

	fmt.Printf("%s\n", finalJson)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
