package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://gist.github.com/lillylangtree/b55828fa05ed3470d352"

func main() {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	bytes, err := io.ReadAll(response.Body)

	content := string(bytes)
	fmt.Println("This is html:\n", content)
}
