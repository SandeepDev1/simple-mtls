package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load the root CA certificate
	caCert, err := os.ReadFile("ca/ca.crt")
	if err != nil {
		log.Fatalf("Failed to read CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatalf("Failed to add CA certificate to pool")
	}

	// Set up the TLS configuration from the cert pool
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	// Create the HTTP client with the TLS Configuration
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// Make a request to the http server which has mTLS
	resp, err := client.Get("https://localhost:8443")
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}
	log.Printf("Response: %s", body)
}
