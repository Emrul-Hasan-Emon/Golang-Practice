package main

import (
	"fmt"
	"time"
)

func workerPool(numOfWorkers int, jobs <-chan string, results chan<- string) {
	for w := 1; w <= numOfWorkers; w++ {
		go worker(w, jobs, results)
	}
}

func worker(worker int, jobs <-chan string, results chan<- string) {
	for url := range jobs {
		fmt.Printf("Worker %d started fetching %s\n", worker, url)
		results <- process(url)
		fmt.Printf("Worker %d completed fetching %s\n", worker, url)
	}
}

func process(url string) string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("Web scrapping done for Url %s", url)
}

func main() {
	// Web pages from where data will be fetched
	urls := []string{
		"http://example.com",
		"http://example.org",
		"http://example.net",
		"http://example.com/about",
		"http://example.org/c",
		"http://example.org/contact",
		"http://example.org/b",
		"http://example.org/a",
	}
	numOfWorkers := 3
	jobs := make(chan string, len(urls))
	results := make(chan string, len(urls))

	// Start the worker pool
	workerPool(numOfWorkers, jobs, results)

	// Send URLs to job queue
	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	// Collect results
	for a := 1; a <= len(urls); a++ {
		result := <-results
		fmt.Println("Result: ", result)
	}
	fmt.Println("Done..............")
}
