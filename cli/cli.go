package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/dev-t2/learn-cryptocurrency/api"
	"github.com/dev-t2/learn-cryptocurrency/explorer"
)

func Start() {
	mode := flag.String("mode", "api", "")
	port := flag.Int("port", 8080, "")

	flag.Parse()

	switch *mode {
	case "api":
		api.Start(*port)
	case "explorer":
		explorer.Start(*port)
	default:
		fmt.Println("Please enter the following flags")
		fmt.Println("-mode=api (api or explorer)")
		fmt.Println("-port=8080")

		os.Exit(0)
	}
}