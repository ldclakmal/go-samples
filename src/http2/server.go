// How to run the program
// $ go run server.go
// $ curl --http2 -X POST http://localhost:9090 -d "Hello"

package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"net/http"
)

func main() {
	var srv http.Server
	srv.Addr = ":9090"
	//Enable http2
	http2.ConfigureServer(&srv, nil)
	http.HandleFunc("/", echo)
	srv.ListenAndServe()
}

func echo(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	contents, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Fprintf(w, "%s\n", string(contents))
}
