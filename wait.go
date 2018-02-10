package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	wait := new(sync.WaitGroup)
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}

	for _, url := range urls {

		// waitgroupに追加
		wait.Add(1)

		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()

			fmt.Println(url, res.Status)

			// waitgroupから削除
			wait.Done()
		}(url)
	}

	// 待ち合わせ
	wait.Wait()
}
