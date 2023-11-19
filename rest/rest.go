package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dev-t2/learn-cryptocurrency/blockchain"
	"github.com/dev-t2/learn-cryptocurrency/utils"
)

var addr string

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", addr, u)

	return []byte(url), nil
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type addBlock struct {
	Data string
}

func docs(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	
	data := []urlDescription{
		{ 
			URL: url("/"), 
			Method: "GET", 
			Description: "Documentation",
		},
		{ 
			URL: url("/blocks"), 
			Method: "GET", 
			Description: "Blocks",
		},
		{ 
			URL: url("/blocks"), 
			Method: "POST", 
			Description: "Add Block", 
			Payload: "data: string", 
		},
		{ 
			URL: url("/blocks/{id}"), 
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
		var addBlock addBlock

		utils.HandleErr(json.NewDecoder(req.Body).Decode(&addBlock)) 

		blockchain.GetBlockchain().AddBlock(addBlock.Data)

		res.WriteHeader(http.StatusCreated)
	}
}

func Start(port int) {
	handler := http.NewServeMux()

	handler.HandleFunc("/", docs)

	handler.HandleFunc("/blocks", blocks)

	addr = fmt.Sprintf(":%d", port)

	fmt.Printf("Listening on http://localhost%s\n", addr)

	log.Fatal(http.ListenAndServe(addr, handler))
}