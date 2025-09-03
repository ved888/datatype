package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func GetData(url string) interface{} {
	res, err := http.Get("https://" + url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return (string(body))
	// return res.Status
}

func main() {
	urls := []string{"www.google.com", "www.facebook.com"}
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			data := GetData(url)
			fmt.Println("Data :", data)
		}(url)
	}
	wg.Wait()
}
