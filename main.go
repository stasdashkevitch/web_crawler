package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 4 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: crawler <baseURL> <maxConcurrency> <maxPages>")
		return
	}

	if len(args) > 4 {
		fmt.Println("too many arguments provided")
		return
	}

	rawBaseURl := os.Args[1]
	maxConcurrencyString := os.Args[2]
	maxPagesString := os.Args[3]

	maxConcurrency, err := strconv.Atoi(maxConcurrencyString)
	if err != nil {
		fmt.Printf("Error - maxConcurrency: %v", err)
	}

	maxPages, err := strconv.Atoi(maxPagesString)
	if err != nil {
		fmt.Printf("Error - maxPages: %v", err)
	}

	cfg, err := configure(rawBaseURl, maxConcurrency, maxPages)
	if err != nil {
		log.Fatal("Error - configure: %v", err)
	}

	fmt.Printf("starting crawling of: %s ... \n", rawBaseURl)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURl)
	cfg.wg.Wait()

	printReport(cfg.pages, rawBaseURl)

}
