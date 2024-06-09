package main

import (
	"fmt"
	"net/http"

	"github.com/zeeshanahmad0201/learning_go/tree/main/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	fmt.Println("Server is getting started...")

	r := router.Router()
	http.ListenAndServe(":8000", r)
	fmt.Println("Listening at port 8000...")

}
