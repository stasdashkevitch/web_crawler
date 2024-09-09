package main

import (
	"log"
	"net/url"
	"strings"
)

func NormalizeUrl(inputURL string) (string, error) {
	URL, err := url.Parse(inputURL)
	if err != nil {
		log.Fatal("Cannot parse url")
		return "", err
	}

	norm_url := URL.Host + URL.Path
	norm_url = strings.ToLower(norm_url)
	norm_url = strings.TrimSuffix(norm_url, "/")

	return norm_url, nil
}
