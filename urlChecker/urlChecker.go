package urlChecker

import (
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

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
	c := make(chan result)

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		hitURLResult := <-c
		results[hitURLResult.url] = hitURLResult.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- result) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- result{
		url:    url,
		status: status,
	}
}
