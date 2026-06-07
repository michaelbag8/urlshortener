package main

import (
	"fmt"
	"log"
	"net/http"
)
var urlStore = map[string]URL{}

func main() {
	http.HandleFunc("GET /health", healthHandler)
	http.HandleFunc("POST /urls", createURLHandler)
	http.HandleFunc("GET /{code}", redirectHandler)

	fmt.Println("server is running at http://localhost:8080/")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
