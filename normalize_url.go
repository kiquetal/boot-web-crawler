package main

import (
	"fmt"
	"net/url"
	"errors"
)

func normalizeURL(url string) (string, error) {
	u, err := url.Parse("https://example.org")

	if err != nil {
		return errors.New("failed to parse URL")
	}
	fmt.Println(u.Host)
}
