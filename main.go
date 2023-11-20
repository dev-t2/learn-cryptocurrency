package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Please enter the following command.")
	fmt.Println()
	fmt.Println("explorer:  HTML Explorer")
	fmt.Println("api:       REST API (Recommended)")
	fmt.Println()

	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start HTML Explorer")
	case "api":
		fmt.Println("Start REST API")
	default:
		usage()
	}
}