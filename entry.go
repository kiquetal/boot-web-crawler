package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kiquetal/boot-web-crawler/internal/database"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	// check if the executable have received 1 parameter
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error opening db")
		os.Exit(1)
	}
	fmt.Printf("db: %v\n", db)
	fmt.Printf("DB URL: %v\n", dbURL)
	defer db.Close()

	_ = database.New(db)

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
