// How to run the program
// $ go run http_small_backend.go
// $ curl -X POST http://localhost:9191/hello/sayHello -d "Hello"

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello/sayHello", smallPayload)
	log.Printf("Serving on http://localhost:9191/hello/sayHello")
	log.Fatal(http.ListenAndServe(":9191", nil))
}

func smallPayload(w http.ResponseWriter, req *http.Request) {
	//log.Printf("Request connection: %s, path: %s", req.Proto, req.URL.Path[1:])
	payload := "{ \"foo\": \"bar\" }"
	fmt.Fprintf(w, "%s\n", payload)
}
