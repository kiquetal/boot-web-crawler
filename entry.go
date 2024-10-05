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
	// get the html body from the url
	htmlBody, err := getHTML(url)
	if err != nil {
		fmt.Println("error getting html")
		os.Exit(1)
	}
	fmt.Println("html body retrieved")
	fmt.Println("body length: ", len(htmlBody))
	fmt.Println("body content: ", htmlBody)

}
