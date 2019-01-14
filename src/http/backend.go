// How to run the program
//
// Run as HTTP/1.1 echo backend
// $ go run backend.go
// Test - $ curl -k https://localhost:9191/hello/sayHello -d "Hello"
//
// Run as HTTP/2 echo backend
// $ go run backend.go -v 2
// Test - $ curl --http2 -k https://localhost:9191/hello/sayHello -d "Hello"

package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
	"net/http"
)

// By default version flag is set to 1 (refers to HTTP/1.1)
var httpVersion = flag.Int("v", 1, "HTTP version")

func main() {
	flag.Parse()

	switch *httpVersion {
	case 1:
		httpBackend()
	case 2:
		http2Backend()
	}
}

func httpBackend() {
	http.HandleFunc("/hello/sayHello", echoPayload)
	log.Printf("Serving HTTP/1.1 on https://localhost:9191/hello/sayHello")
	log.Fatal(http.ListenAndServeTLS(":9191", "../cert/server.crt", "../cert/server.key", nil))
}

func http2Backend() {
	var httpServer = http.Server{
		Addr: ":9191",
	}
	_ = http2.ConfigureServer(&httpServer, nil)
	http.HandleFunc("/hello/sayHello", echoPayload)
	log.Printf("Serving HTTP/2 on https://localhost:9191/hello/sayHello")
	log.Fatal(httpServer.ListenAndServeTLS("../cert/server.crt", "../cert/server.key"))
}

func echoPayload(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request connection: %s, path: %s", req.Proto, req.URL.Path[1:])
	defer req.Body.Close()
	contents, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Oops! Failed reading body of the request.\n %s", err)
		http.Error(w, err.Error(), 500)
	}
	fmt.Fprintf(w, "%s\n", string(contents))
}
