package main

import (
	"fmt"
	"net/http"
	"log"
)

func getStaus(urls []string) <-chan string {
	statusChan := make(chan string)

	for _,url := range urls {
		go func(url string) {
			res,err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			statusChan <- res.Status
		}(url)
	}

	return statusChan
}

func main() {
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}
	statusChan := getStaus(urls)

	for i:=0; i<len(urls);i++  {
	fmt.Println(<-statusChan)
	}
}
