package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/dev-t2/learn-cryptocurrency/01-blockchain/blockchain"
)

type homeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}

const (
	port = ":8080"
	dir = "02-explorer/templates"
)

var templates *template.Template

func homeHandler (res http.ResponseWriter, req *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}

	templates.ExecuteTemplate(res, "home", data)
}

func addHandler (res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		templates.ExecuteTemplate(res, "add", nil)
	case "POST":
		req.ParseForm()

		data := req.Form.Get("data")

		blockchain.GetBlockchain().AddBlock(data)

		http.Redirect(res, req, "/", http.StatusPermanentRedirect)
	}
}

func main() {
	templates = template.Must(template.ParseGlob(dir + "/pages/*.html"))
	templates = template.Must(templates.ParseGlob(dir + "/partials/*.html"))

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/add", addHandler)

	fmt.Printf("Listening on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}