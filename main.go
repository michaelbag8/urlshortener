package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("GET /health", handleHealth)
	http.HandleFunc("POST /urls", createURLHandler)
	http.HandleFunc("GET /{code}", redirectHandler)

	fmt.Println("server is running at http://localhost:8080/")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("failed to start server")
		return
	}

}
