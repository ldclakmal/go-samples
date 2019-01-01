package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

const URL = "http://localhost:9191/hello/sayHello"

func main() {
	resp, err := http.Post(URL, "text/plain", bytes.NewBufferString("Hello"))
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Response body: %s", string(body))
}
