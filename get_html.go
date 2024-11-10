package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
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

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func (c *config) crawlPage(rawCurrentlURL string) {
	// get the html body from the url

	if len(c.pages) >= c.maxPages {
		return
	}

	c.concurrencyControl <- struct{}{} // acquire a slot
	c.wg.Add(1)                        // increment the WaitGroup counter

	go func() {
		defer c.wg.Done() // decrement the WaitGroup counter when done
		defer func() {
			<-c.concurrencyControl

			fmt.Println("slot released")
		}() // release the slot

		// get the html body from the u

		if !strings.HasPrefix(rawCurrentlURL, "http") {
			rawCurrentlURL = c.baseURL.Scheme + "://" + c.baseURL.Host + rawCurrentlURL
		}
		fmt.Println("current u: ", rawCurrentlURL)
		htmlBody, err := getHTML(rawCurrentlURL)
		if err != nil {
			fmt.Println("error getting html")
			return
		}
		fmt.Println("html body retrieved")
		fmt.Println("body length: ", len(htmlBody))
		fmt.Println("body content: ", htmlBody)
		urls, err := getURLFromHTML(rawCurrentlURL, htmlBody)
		if err != nil {
			fmt.Println("error getting urls")
			return
		}
		fmt.Println("urls retrieved")
		fmt.Println("urls: ", urls)
		for _, u := range urls {
			normalizedURL, err := normalizeURL(u)
			// check if the domain is the same as the base u

			if strings.HasPrefix(normalizedURL, "/") {
				fmt.Println("same domain")
			} else {
				parsedURL, err := url.Parse(normalizedURL)
				if err != nil {
					fmt.Println("Error parsing URL:", err)
					return
				}
				if parsedURL.Scheme == "" || parsedURL.Host == "" {
					fmt.Println("same domain")
				} else {
					fmt.Println("different domain")
					fmt.Println("normalized URL:", normalizedURL)
					continue
				}
			}
			if err != nil {
				fmt.Println("error normalizing u")
				continue
			}
			if _, ok := c.pages[normalizedURL]; ok {
				c.mu.Lock()
				c.pages[normalizedURL]++
				c.mu.Unlock()
			} else {
				c.mu.Lock()
				c.pages[normalizedURL] = 1
				c.mu.Unlock()
				c.crawlPage(normalizedURL)
			}

		}

	}()
}
