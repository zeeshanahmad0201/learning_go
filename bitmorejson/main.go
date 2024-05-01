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

	// EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsonFromWeb := []byte(`
	{
		"title": "flutter development",
		"price": 199,
		"tags": [
				"flutter",
				"dev"
		]
	}`)
	if json.Valid(jsonFromWeb) {
		fmt.Println("json is valid")
		var lcoCourse course
		json.Unmarshal(jsonFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
		fmt.Printf("%s\n", lcoCourse.Tags)
	} else {
		fmt.Println("JSON IS NOT VALID")
	}

	var onlineData map[string]interface{}
	json.Unmarshal(jsonFromWeb, &onlineData)

	for k, v := range onlineData {
		fmt.Printf("Key is %v and Value is %v\n", k, v)
	}
}
