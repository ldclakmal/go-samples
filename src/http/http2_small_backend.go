// How to run the program
// $ go run http2_small_backend.go
// $ curl --http2 -X POST http://localhost:9191/nyseStock/stocks -d "Hello"

package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"log"
	"net/http"
)

func main() {
	var srv http.Server
	srv.Addr = ":9191"
	//Enable http2
	http2.ConfigureServer(&srv, nil)
	http.HandleFunc("/nyseStock/stocks", small_payload)
	log.Printf("Serving on http://localhost:9191/nyseStock/stocks")
	log.Fatal(srv.ListenAndServeTLS("../cert/server.crt", "../cert/server.key"))
}

func small_payload(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request connection: %s, path: %s", req.Proto, req.URL.Path[1:])
	payload := "{ \"foo\": \"bar\" }"
	fmt.Fprintf(w, "%s\n", payload)
}
