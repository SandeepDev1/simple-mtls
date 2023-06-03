package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load the server the certificates from the local file
	cert, err := tls.LoadX509KeyPair("server/server.crt", "server/server.key")
	if err != nil {
		log.Fatalf("Failed to load server key pair: %v", err)
	}

	// Create a pool of certificate with our root CA
	// which acts as an authenticator with our clients
	caCert, err := os.ReadFile("ca/ca.crt")
	if err != nil {
		log.Fatalf("Failed to read CA certificate: %v", err)
	}
	caPool := x509.NewCertPool()
	if ok := caPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal("Failed to add CA certificate to pool")
	}

	// Set up the TLS config
	// ClientAuth - VerifyClientCertIfGiven -> This tells to verify the root CA and if the cert is provided by the client verify that cert too or else no need
	// ClientCAs - it is our cert pool from root CA
	// Certificates - it is our server certificate to encrypt and decrypt traffic between the client and server
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.VerifyClientCertIfGiven,
		ClientCAs:    caPool,
	}

	// Create the http server and listen on some port, also pass the tls config
	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	// Create the http handler and handle the requests in that function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, this is the server"))
	})

	// Listen for the connections and server the TLS
	log.Fatal(server.ListenAndServeTLS("server/server.crt", "server/server.key"))
}
