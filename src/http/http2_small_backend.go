// How to run the program
// $ go run http2_small_backend.go
// $ curl --http2 -X POST http://localhost:9090 -d "Hello"

package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"net/http"
)

func main() {
	var srv http.Server
	srv.Addr = ":9191"
	//Enable http2
	http2.ConfigureServer(&srv, nil)
	http.HandleFunc("/nyseStock/stocks", small_payload)
	srv.ListenAndServe()
}

func small_payload(w http.ResponseWriter, req *http.Request) {
	payload := "{ \"foo\": \"bar\" }"
	fmt.Fprintf(w, "%s\n", payload)
}
