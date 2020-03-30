package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	//초기화 방법 1번
	var results = make(map[string]string)
	//초기화 방법 2번 초기화 시켜줘야 쓸수 있다 Map 도 new HashMap 해야 쓸 수 있잖아

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAIL"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking === ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	} else {
		return nil
	}
}
