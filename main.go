package main

import (
	"fmt"
	"net/http"

	"github.com/Shared-Info-Platform/url-query-transformer/handler"
)

func main() {

	http.HandleFunc("/", handler.HandleRequest)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
