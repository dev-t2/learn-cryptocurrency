package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

type URLDescription struct {
	URL         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func documentation(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	
	data := []URLDescription{
		{ URL: "/", Method: "GET", Description: "Documentation" },
		{ URL: "/blocks", Method: "POST", Description: "Add Block", Payload: "data:string" },
	}

	json.NewEncoder(res).Encode(data)
}

func main() {
	http.HandleFunc("/", documentation)

	fmt.Printf("Listening on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}