package main

import (
	"fmt"
	"log"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ok")
}

func main() {
	http.HandleFunc("/", Health)
	log.Fatal(http.ListenAndServe(":9100", nil))
}
