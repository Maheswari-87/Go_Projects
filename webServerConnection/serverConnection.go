package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", greet)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
