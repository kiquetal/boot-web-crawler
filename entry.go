package main

import (
	"fmt"
	"os"
)

func main() {
	// check if the executable have received 1 parameter
	if len(os.Args) < 2 {
		// if not, exit with status code 1
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		// if not, exit with status code 1
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	url := os.Args[1]
	fmt.Println("starting crawl")
	fmt.Printf("%v\n", url)

}