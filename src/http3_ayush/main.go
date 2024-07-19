package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/quic-go/quic-go/http3"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, HTTP/3!\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	// Start the HTTP/3 server with self-signed certificates
	log.Println("Starting HTTP/3 server on :8080")
	err := http3.ListenAndServeQUIC("127.0.0.1:8080", "certs/cert.pem", "certs/key.pem", mux)
	if err != nil {
		log.Fatalf("Failed to start HTTP/3 server: %v", err)
	}
}
