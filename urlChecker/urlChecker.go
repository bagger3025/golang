package urlChecker

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	errRequestFailed = errors.New("Request failed")
)

func UrlCheckerMain() {
	urls := []string{
		"https://www.naver.com",
		"https://www.google.com",
		"https://soundcloud.com",
		"https://www.airbnb.com",
		"https://www.redit.com",
		"https://www.amazon.com",
		"https://academy.nomadcoders.io",
		"https://www.github.com",
		"https://www.notion.so",
	}

	for _, url := range urls {
		hitURL(url)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}

	return nil
}
