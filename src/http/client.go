// How to run the program
// $ go run client.go
// $ go run client.go -d "Hello Go"
// $ go run client.go -v 2

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
	resp, err := client.Post("http://localhost:9090/passthrough", "text/plain", bytes.NewBufferString(*requestBody))
	//resp, err := client.Post("https://localhost:9191/hello/sayHello", "text/plain", bytes.NewBufferString("Hello"))
	//resp, err := client.Get(url)
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
