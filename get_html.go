package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(raw string) (string, error) {
	res, err := http.Get(raw)
	if err != nil {
		return "", fmt.Errorf("got network error: %v", err)
	}

	if res.StatusCode > 399 {
		return "", fmt.Errorf("got HTTP error: %v", err, res.Status)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("got none HTML response: %v", contentType)
	}

	htmlBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body", err)
	}

	htmlBody := string(htmlBodyBytes)

	return htmlBody, nil
}
