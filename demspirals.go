package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"hello\":\"world\"}")
}
