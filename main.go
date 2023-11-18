package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dev-t2/learn-cryptocurrency/utils"
)

const port = ":8080"

type URLDescription struct {
	URL         string
	Method      string
	Description string
}

func documentation(res http.ResponseWriter, req *http.Request) {
	data := []URLDescription{
		{ URL: "/", Method: "GET", Description: "Get Documentation" },
	}

	b, err := json.Marshal(data)

	utils.HandleErr(err)

	fmt.Printf("%s\n", b)
}

func main() {
	http.HandleFunc("/", documentation)

	fmt.Printf("Listening on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}