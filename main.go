package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dev-t2/learn-cryptocurrency/blockchain"
	"github.com/dev-t2/learn-cryptocurrency/utils"
)

const port = ":8080"

type URL string

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)

	return []byte(url), nil
}

type URLDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type AddBlock struct {
	Data string
}

func docs(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	
	data := []URLDescription{
		{ 
			URL: URL("/"), 
			Method: "GET", 
			Description: "Documentation",
		},
		{ 
			URL: URL("/blocks"), 
			Method: "GET", 
			Description: "Blocks",
		},
		{ 
			URL: URL("/blocks"), 
			Method: "POST", 
			Description: "Add Block", 
			Payload: "data: string", 
		},
		{ 
			URL: URL("/blocks/{id}"), 
			Method: "GET", 
			Description: "Block",
		},
	}

	json.NewEncoder(res).Encode(data)
}

func blocks(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		res.Header().Add("Content-Type", "application/json")

		json.NewEncoder(res).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlock AddBlock

		utils.HandleErr(json.NewDecoder(req.Body).Decode(&addBlock)) 

		blockchain.GetBlockchain().AddBlock(addBlock.Data)

		res.WriteHeader(http.StatusCreated)
	}
}

func main() {
	http.HandleFunc("/", docs)

	http.HandleFunc("/blocks", blocks)

	fmt.Printf("Listening on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}