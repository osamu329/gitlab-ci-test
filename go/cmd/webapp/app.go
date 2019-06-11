package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, rqst *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}
