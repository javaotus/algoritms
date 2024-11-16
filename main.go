package main

import (
	"fmt"
	"sync"
)

func main() {

	url := BaseUrl

	numPages := 10

	var wg sync.WaitGroup
	results := make(chan string)

	for page := 1; page <= numPages; page++ {
		wg.Add(1)
		go processPage(url, page, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}

}
