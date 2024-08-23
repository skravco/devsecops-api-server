package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/success", handleSuccess)
	http.HandleFunc("/health", handleHealth)

	addr := "localhost:8000"
	log.Printf("Listening on %s ...", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleSuccess(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("success Endpoint Triggered")
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("health Endpoint Triggered")
}
