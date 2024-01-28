package main

import (
	"log"
	"sync"
	"unicode/utf8"
)

func fetchWebsitesConcurrently(websites []Website) map[string]string {
	results := make(map[string]string)
	var wg sync.WaitGroup
	resultChan := make(chan struct {
		URL     string
		Content string
	})

	for _, w := range websites {
		wg.Add(1)
		go func(website Website) {
			log.Printf("Handling %s", website.URL)
			defer wg.Done()
			content, err := FetchWebsite(website)
			if err == nil {
				resultChan <- struct {
					URL     string
					Content string
				}{website.URL, content}
			}
		}(w)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		results[result.URL] = result.Content
	}

	return results
}

func main() {
	websites := []Website{
		{"https://nos.nl"},
		{"https://nu.nl"},
		{"http://www.telegraaf.nl"},
		// Add more websites as needed
	}

	results := fetchWebsitesConcurrently(websites)

	for url, content := range results {
		log.Printf("Website: %s, Content: %d ", url, utf8.RuneCountInString(content))
	}
}
