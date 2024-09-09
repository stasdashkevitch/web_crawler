package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {

	list_of_url := []string{
		"https://blog.boot.dev/path/",
		"https://blog.boot.dev/path/",
		"http://blog.boot.dev/path/",
		"http://blog.boot.dev/path",
	}

	expect := "blog.boot.dev/path"

	for _, url := range list_of_url {
		t.Run(fmt.Sprintf("normalize url: %s", url), func(t *testing.T) {
			output, err := NormalizeUrl(url)
			if err != nil {
				t.Errorf("[FAIl]: unexpected error: %v", err)
				return
			}
			if output != expect {
				t.Errorf("[FAIL]: unvalid normalize url, expected: %s, output: %s", expect, output)
				return
			}

		})

	}

}
