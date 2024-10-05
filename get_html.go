package main

import (
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {

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
