package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dev-t2/learn-cryptocurrency/blockchain"
	"github.com/dev-t2/learn-cryptocurrency/utils"
	"github.com/gorilla/mux"
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

func block(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	fmt.Println(id)
}

func Start(port int) {
	router := mux.NewRouter()

	router.HandleFunc("/", docs).Methods("GET")

	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")

	router.HandleFunc("/blocks/{id:[0-9]+}", block).Methods("GET")

	addr = fmt.Sprintf(":%d", port)

	fmt.Printf("Listening on http://localhost%s\n", addr)

	log.Fatal(http.ListenAndServe(addr, router))
}