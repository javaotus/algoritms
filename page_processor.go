package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func processPage(url string, page int, results chan<- string, wg *sync.WaitGroup) {

	defer wg.Done()

	fullURL := fmt.Sprintf("%s/?currentPage=%d", url, page)

	resp, err := http.Get(fullURL)

	if err != nil {
		log.Printf("Error page requesting ... %d: %v", page, err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("Unable to fetch a page %d, error code: %d", page, resp.StatusCode)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Printf("Error parcing page %d: %v", page, err)
		return
	}

	doc.Find(MainContainerSection).Each(func(i int, s *goquery.Selection) {
		s.Find(DivTag).Each(func(j int, inner *goquery.Selection) {
			inner.Find(MetaTag).Each(func(k int, meta *goquery.Selection) {
				if productName, exists := meta.Attr(ProductSectionDiv); exists {
					results <- fmt.Sprintf("Page # %d, product name: %s", page, productName)
				}
			})
		})
	})

}
