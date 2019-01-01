// How to run the program
// $ go run http_echo_backend.go
// $ curl -X POST http://localhost:9191/hello/sayHello -d "Hello"

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello/sayHello", echoHttp)
	log.Printf("Serving on http://localhost:9191/hello/sayHello")
	log.Fatal(http.ListenAndServe(":9191", nil))
}

func echoHttp(w http.ResponseWriter, req *http.Request) {
	//log.Printf("Request connection: %s, path: %s", req.Proto, req.URL.Path[1:])
	defer req.Body.Close()
	contents, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Oops! Failed reading body of the request.\n %s", err)
		http.Error(w, err.Error(), 500)
	}
	fmt.Fprintf(w, "%s\n", string(contents))
}
