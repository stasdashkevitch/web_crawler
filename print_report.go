package main

import (
	"fmt"
	"sort"
)

type Page struct {
	URL   string
	Count int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
=============================
  REPORT for %s
=============================
`, baseURL)

	sortedPages := sortPage(pages)

	for _, page := range sortedPages {
		url := page.URL
		count := page.Count

		fmt.Printf("Found %d internal links to %s\n", count, url)
	}
}

func sortPage(pages map[string]int) []Page {
	pageSlice := []Page{}

	for url, count := range pages {
		pageSlice = append(pageSlice, Page{URL: url, Count: count})
	}

	sort.Slice(pageSlice, func(i, j int) bool {
		if pageSlice[i].Count == pageSlice[j].Count {
			return pageSlice[i].URL < pageSlice[j].URL
		}

		return pageSlice[i].Count > pageSlice[j].Count
	})

	return pageSlice
}
