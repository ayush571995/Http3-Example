package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/quic-go/quic-go/http3"
)

func main() {
	// Load the self-signed certificate
	certBytes, err := ioutil.ReadFile("../certs/cert.pem")
	if err != nil {
		log.Fatalf("Failed to read certificate: %v", err)
	}

	// Create a certificate pool and add the self-signed certificate
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(certBytes)

	// Configure TLS with the certificate pool
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}
	tlsConfig.BuildNameToCertificate()

	// Create an HTTP/3 client with the custom TLS configuration
	client := &http.Client{
		Transport: &http3.RoundTripper{
			TLSClientConfig: tlsConfig,
		},
	}

	// Send a request to the HTTP/3 server
	resp, err := client.Get("https://127.0.0.1:8080")
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read and print the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Printf("Response: %s\n", body)
}
