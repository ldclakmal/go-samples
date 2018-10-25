// How to run the program
// $ go run http2_echo_backend.go
// $ curl --http2 -X POST http://localhost:9090/nyseStock/stocks -d "Hello"

package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"net/http"
)

func main() {
	var srv http.Server
	srv.Addr = ":9191"
	//Enable http2
	http2.ConfigureServer(&srv, nil)
	http.HandleFunc("/nyseStock/stocks", echo)
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
