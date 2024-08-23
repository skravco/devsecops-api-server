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
	// fmt.Println("success Endpoint Triggered")
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	writeResponse(writer, "success!")
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	//fmt.Println("health Endpoint Triggered")
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	writeResponse(writer, "ok!")
}

func handleBrandNewEndpoint(writer http.ResponseWriter, request *http.Request) {
	//fmt.Println("brand-new Endpoint Triggered")
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	writeResponse(writer, "brand-new-endpoint!")
}

func writeResponse(writer http.ResponseWriter, responseString string) {
	response := []byte(responseString)
	fmt.Println(response)
	_, err := writer.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}
