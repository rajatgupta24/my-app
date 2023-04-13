package main

import (
	"fmt"
	"log"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ok")
	fmt.Fprintf(w, "Ok")
}

func main() {
	http.HandleFunc("/", health)
	log.Fatal(http.ListenAndServe(":9100", nil))
}
