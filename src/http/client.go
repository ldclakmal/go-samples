// How to run the program
//
// Run as HTTP/1.1 client
// $ go run client.go -url "http://localhost:9090/passthrough" -d "Hello Go"

// Run as HTTP/2 client
// $ go run client.go -v 2 -url "https://localhost:9090/passthrough" -d "Hello Go"

package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

// By default version flag is set to 1 (refers to HTTP/1.1)
var httpVersion = flag.Int("v", 1, "HTTP version")
var requestUrl = flag.String("url", "http://localhost:9191/hello/sayHello", "Request URL")
var requestBody = flag.String("d", "Hello", "Request body")

func main() {
	flag.Parse()
	client := &http.Client{}

	// Create a pool with the server certificate since it is not signed
	// by a known CA
	caCert, err := ioutil.ReadFile("../cert/server.crt")
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS configuration with the certificate of the server
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	// Use the proper transport in the client
	switch *httpVersion {
	case 1:
		client.Transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	case 2:
		client.Transport = &http2.Transport{
			TLSClientConfig: tlsConfig,
		}
	}

	// Perform the request
	resp, err := client.Post(*requestUrl, "text/plain", bytes.NewBufferString(*requestBody))
	if err != nil {
		log.Fatalf("Failed get: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed reading response body: %s", err)
	}
	fmt.Printf("Got response %d: %s %s", resp.StatusCode, resp.Proto, string(body))
}
