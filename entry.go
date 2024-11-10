package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kiquetal/boot-web-crawler/internal/database"
	_ "github.com/lib/pq"
	urlLib "net/url"
	"os"
	"sort"
	"strconv"
	"sync"
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
	if len(os.Args) > 4 {
		// if not, exit with status code 1
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	url := os.Args[1]
	maxPages, _ := strconv.Atoi(os.Args[2])
	maxConcurrency, _ := strconv.Atoi(os.Args[3])
	fmt.Println("starting crawl")
	fmt.Printf("%v\n", url)
	// get the html body from the url
	//	htmlBody, err := getHTML(url)
	baseURL, _ := urlLib.Parse(url)
	c := config{
		pages:              make(map[string]int),
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		baseURL:            baseURL,
		maxPages:           maxPages,
	}

	c.crawlPage(url)
	c.wg.Wait()
	printReport(c.pages, url)
	if err != nil {
		fmt.Println("error getting html")
		os.Exit(1)
	}

}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("Length of pages: ", len(pages))
	fmt.Printf("\n========================================\n")
	fmt.Printf("REPORT for %v\n", baseURL)
	fmt.Printf("========================================\n")

	//sort the map by value

	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range pages {
		ss = append(ss, kv{k, v})
	}

	// Sort the slice based on the map values
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	// Print the sorted key-value pairs
	for _, kv := range ss {
		fmt.Printf("Found %v internal links to %v\n", kv.Value, kv.Key)

	}

}
