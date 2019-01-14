package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

var client = &http.Client{}

func main() {
	client.Transport = &http.Transport{}
	http.HandleFunc("/passthrough", forwardProxy)
	log.Printf("Serving passthrough on http://localhost:9090/passthrough")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func forwardProxy(w http.ResponseWriter, req *http.Request) {
	backendReq, _ := http.NewRequest(req.Method, "http://localhost:9191/hello/sayHello", req.Body)
	resp, _ := client.Do(backendReq)
	body, _ := ioutil.ReadAll(resp.Body)
	_, _ = w.Write(body)
}
