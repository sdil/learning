package main

import (
	"fmt"
	"sync"
	"net/http"
)

var urls = []string{
	"https://fadhil-blog.dev",
	"https://google.com",
	"https://yahoo.com",
	"https://youtube.com",
	"https://amazon.com",
	"https://shopee.com",
	"https://uber.com",
}

func fetch(url string, wg *sync.WaitGroup) {
	fmt.Printf("requesting to url: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	wg.Done()
	fmt.Printf("url: %s status: %s\n", url, resp.Status)
}

func main() {
	fmt.Println("Get all Urls")

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
	}

	wg.Wait()
	fmt.Println("All done")
}