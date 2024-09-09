package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

var ErrInvalidHTML = errors.New("cannot parse HTML body")
var ErrInvalidBaseURL = errors.New("invalid base URL")
var ErrInvalidHrefURL = errors.New("invalid href URL")

func getUrlFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	base, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		fmt.Errorf("[ERROR]: cannot parse HTML body")
		return nil, ErrInvalidHTML
	}

	var result []string

	var traversNode func(*html.Node) error
	traversNode = func(n *html.Node) error {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href, err := url.Parse(a.Val)
					if err != nil {
						fmt.Errorf("[ERROR]: cannot parse href URL")
						return ErrInvalidHrefURL
					}
					resolvedURL := baseURL.ResolveReference(href)
					result = append(result, resolvedURL.String())
					break
				}
			}
		}

		for child := n.FirstChild; child != nil; child = child.NextSibling {
			traversNode(child)
		}

		return nil
	}

	err = traversNode(base)
	if err != nil {
		return nil, ErrInvalidHrefURL
	}

	return result, nil
}
