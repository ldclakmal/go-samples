// How to run the program
// $ go run http2_small_backend.go
// $ curl --http2 -X POST http://localhost:9191/hello/sayHello -d "Hello"

package http2

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
	http.HandleFunc("/hello/sayHello", smallPayload)
	log.Printf("Serving on http://localhost:9191/hello/sayHello")
	log.Fatal(srv.ListenAndServeTLS("../cert/server.crt", "../cert/server.key"))
}

func smallPayload(w http.ResponseWriter, req *http.Request) {
	//log.Printf("Request connection: %s, path: %s", req.Proto, req.URL.Path[1:])
	payload := "{ \"foo\": \"bar\" }"
	fmt.Fprintf(w, "%s\n", payload)
}
