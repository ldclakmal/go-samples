// How to run the program
// $ go run backend.go
// $ curl http://localhost:9191/hello/sayHello -d "Hello"
// $ curl --http2 -k https://localhost:9191/hello/sayHello -d "Hello"

package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	httpServer()
	//http2Server()
}

func httpServer() {
	http.HandleFunc("/hello/sayHello", echoPayload)
	log.Printf("Serving HTTP/1.1 on http://localhost:9191/hello/sayHello")
	log.Fatal(http.ListenAndServe(":9191", nil))
}

func http2Server() {
	var srv http.Server
	srv.Addr = ":9191"
	_ = http2.ConfigureServer(&srv, nil)
	http.HandleFunc("/hello/sayHello", echoPayload)
	log.Printf("Serving HTTP/2 on https://localhost:9191/hello/sayHello")
	log.Fatal(srv.ListenAndServeTLS("../cert/server.crt", "../cert/server.key"))
}

func echoPayload(w http.ResponseWriter, req *http.Request) {
	//log.Printf("Request connection: %s, path: %s", req.Proto, req.URL.Path[1:])
	defer req.Body.Close()
	contents, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Oops! Failed reading body of the request.\n %s", err)
		http.Error(w, err.Error(), 500)
	}
	fmt.Fprintf(w, "%s\n", string(contents))
}

func smallPayload(w http.ResponseWriter, req *http.Request) {
	//log.Printf("Request connection: %s, path: %s", req.Proto, req.URL.Path[1:])
	payload := "Hello from GO HTTP/2"
	fmt.Fprintf(w, "%s\n", payload)
}
