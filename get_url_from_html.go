package main

import (
	"golang.org/x/net/html"
	"strings"
)

func getURLFromHTML(rawBase, htmlBody string) ([]string, error) {

	//find all the a elements in html body
	htmlBodyReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlBodyReader)
	if err != nil {
		return nil, err
	}
	var urls []string
	//search all a link in the html body
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			// get the href attribute
			for _, a := range n.Attr {
				if a.Key == "href" {
					urls = append(urls, a.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}

	}
	f(doc)
	return urls, nil
}
