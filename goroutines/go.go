package main

import (
	"fmt"
	"net/http"
)

func main() {

	c := make(chan result)
	urls := [5]string{
		"https://www.naver.com",
		"https://bigfinance.co.kr",
		"https://www.google.com",
		"https://youtube.com",
		"https://github.com/shmoon0814",
	}

	for _, url := range urls {
		go checkURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

}

func checkURL(url string, c chan result) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- result{url: url, status: status}
}

type result struct {
	url    string
	status string
}
