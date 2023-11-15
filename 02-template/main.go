package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/dev-t2/learn-cryptocurrency/02-template/blockchain"
)

type homeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}

const port = ":8080"

func home (w http.ResponseWriter, r *http.Request) {
	html := template.Must(template.ParseFiles("02-template/templates/pages/home.html"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}

	html.Execute(w, data)
}

func main() {
	http.HandleFunc("/", home)

	fmt.Printf("Listening on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}