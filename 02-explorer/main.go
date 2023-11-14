package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

func home (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func main() {
	http.HandleFunc("/", home)

	fmt.Printf("Listening on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}