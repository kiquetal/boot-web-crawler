package main

import (
	"errors"
	"net/url"
)

func normalizeURL(urlString string) (string, error) {
	u, err := url.Parse(urlString)

	if err != nil {
		return "", errors.New("failed to parse URL")
	}

	//return everything except schema using the Host and Path

	toReturn := u.Host + u.Path
	if toReturn[len(toReturn)-1] == '/' {
		toReturn = toReturn[:len(toReturn)-1]
	}

	return toReturn, nil
}
