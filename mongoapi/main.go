package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zeeshanahmad0201/learning_go/tree/main/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")

	r := router.Router()
	fmt.Println("Server is getting started...")

	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Listening at port 8000...")

}
