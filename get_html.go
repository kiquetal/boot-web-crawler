package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	fmt.Printf("getting html from %v\n", rawURL)
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	status := resp.StatusCode
	if status != 200 {
		return "", err
	}
	body, e := io.ReadAll(resp.Body)
	if e != nil {
		return "", e
	}

	return string(body), nil

}

func crawlPage(rawBaseURL, rawCurrentlURL string, pages map[string]int) {
	// get the html body from the url
	htmlBody, err := getHTML(rawCurrentlURL)
	if err != nil {
		fmt.Println("error getting html")
		return
	}
	fmt.Println("html body retrieved")
	fmt.Println("body length: ", len(htmlBody))
	fmt.Println("body content: ", htmlBody)
	urls, err := getURLFromHTML(rawBaseURL, htmlBody)
	if err != nil {
		fmt.Println("error getting urls")
		return
	}
	fmt.Println("urls retrieved")
	fmt.Println("urls: ", urls)
	for _, url := range urls {
		normalizedURL, err := normalizeURL(url)
		// check if the domain is the same as the base url

		if strings.HasPrefix(normalizedURL, "/") {
			fmt.Println("same domain")
		} else {
			fmt.Println("different domain")
			fmt.Println("normalized url: ", normalizedURL)
			continue
		}

		if err != nil {
			fmt.Println("error normalizing url")
			continue
		}
		if _, ok := pages[normalizedURL]; !ok {
			pages[normalizedURL] = 1
			crawlPage(rawBaseURL, normalizedURL, pages)
		}
	}
	fmt.Println("pages: ", pages)
}
