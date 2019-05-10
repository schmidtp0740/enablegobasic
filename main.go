package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	http.HandleFunc("/", handler)
	fmt.Printf("Listening on port: %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
