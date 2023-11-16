package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/dev-t2/learn-cryptocurrency/blockchain"
)

type homeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}

const (
	port = ":8080"
	dir = "explorer/templates"
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

func Start() {
	templates = template.Must(template.ParseGlob(dir + "/pages/*.html"))
	templates = template.Must(templates.ParseGlob(dir + "/partials/*.html"))

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/add", addHandler)

	fmt.Printf("Listening on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}