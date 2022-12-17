package urlChecker

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	errRequestFailed = errors.New("request failed")
)

func UrlCheckerMain() {
	urls := []string{
		"https://www.naver.com",
		"https://www.google.com",
		"https://soundcloud.com",
		"https://www.airbnb.com",
		"https://www.reddit.com",
		"https://www.amazon.com",
		"https://academy.nomadcoders.co",
		"https://www.github.com",
		"https://www.notion.so",
	}

	results := make(map[string]string)

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}

		results[url] = result
	}

	// fmt.Println(results)

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err)
		fmt.Println(resp.StatusCode)
		return errRequestFailed
	}

	return nil
}
