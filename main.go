package main

import (
	"github.com/dev-t2/learn-cryptocurrency/api"
	"github.com/dev-t2/learn-cryptocurrency/explorer"
)

func main() {
	go api.Start(8080)

	explorer.Start(8081)
}