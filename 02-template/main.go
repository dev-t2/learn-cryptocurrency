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

const (
	port = ":8080"
	dir = "02-template/templates"
)

var templates *template.Template

func home (w http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}

	templates.ExecuteTemplate(w, "home", data)
}

func main() {
	templates = template.Must(template.ParseGlob(dir + "/pages/*.html"))
	templates = template.Must(templates.ParseGlob(dir + "/partials/*.html"))

	http.HandleFunc("/", home)

	fmt.Printf("Listening on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}