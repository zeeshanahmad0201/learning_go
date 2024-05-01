package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const baseUrl = "http://localhost:8000"

func main() {
	fmt.Println("Wecome to get_requests project!")
	// PerformGetRequest(baseUrl)

	PerformPostJsonRequest(baseUrl)
}

func PerformGetRequest(url string) {
	url = url + "/get"

	response, err := http.Get(url)
	checkNilErr(err)

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	// short way
	// content, err := io.ReadAll(response.Body)
	// checkNilErr(err)
	// fmt.Println("Content: ", string(content))

	// using strings package
	var responseString strings.Builder
	contentByte, err := io.ReadAll(response.Body)
	checkNilErr(err)
	responseString.Write(contentByte)

	fmt.Println(responseString.String())

}

func PerformPostJsonRequest(url string) {
	url = url + "/post"

	requestBody := strings.NewReader(`
	{
		"title":"POST request",
		"methodType":"POST"
	}
	`)

	response, err := http.Post(url, "application/json", requestBody)
	checkNilErr(err)

	defer response.Body.Close()

	contentByte, err := io.ReadAll(response.Body)
	checkNilErr(err)

	fmt.Println(string(contentByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
