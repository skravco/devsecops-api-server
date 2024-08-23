package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/success", handleSuccess)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/brand-new-endpoint", handleBrandNewEndpoint)

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

func handleBrandNewEndpoint(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("brand-new Endpoint Triggered")
}
