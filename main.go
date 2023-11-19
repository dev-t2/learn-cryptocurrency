package main

import (
	"github.com/dev-t2/learn-cryptocurrency/explorer"
	"github.com/dev-t2/learn-cryptocurrency/rest"
)

func main() {
	go rest.Start(8080)

	explorer.Start(8081)
}