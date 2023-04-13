package main

import (
	"fmt"
	"log"
	"net/http"
)

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ok")
}

func main() {
	http.HandleFunc("/health", fooHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
