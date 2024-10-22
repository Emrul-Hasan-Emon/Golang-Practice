package main

import (
	"fmt"
	"log"
	"net/http"
)

var requestID int

func initializeRequestId() {
	requestID = 0
}

type Job struct {
	id      int
	request *http.Request
}

func workerPool(numOfWorkers int, jobs <-chan Job, results chan<- string) {
	for w := 1; w <= numOfWorkers; w++ {
		go worker(w, jobs, results)
	}
}

func worker(w int, jobs <-chan Job, results chan<- string) {
	for job := range jobs {
		fmt.Printf("Worker %d started processing request %d\n", w, job.id)
		results <- process(w, job)
		fmt.Printf("Worker %d finished processing request %d\n", w, job.id)
	}
}

func process(w int, job Job) string {
	return fmt.Sprintf("Worker %d processed job with job id %d", w, job.id)
}

func requestHandler(jobs chan<- Job, results <-chan string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID++
		job := Job{id: requestID, request: r}

		// Send the job to the job queue
		jobs <- job

		// Respond to the client
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(fmt.Sprintf("Request %d is being processed\n", requestID)))
	}
}

func resultListener(results <-chan string) {
	for result := range results {
		fmt.Println("Result: ", result)
	}
	fmt.Println("Done....")
}

func main() {
	numOfWorkers := 3
	jobs := make(chan Job)
	results := make(chan string)

	// Start worker pool
	workerPool(numOfWorkers, jobs, results)

	// Initialize Request ID
	initializeRequestId()

	// Setup HTTP server with request handler
	http.HandleFunc("/process", requestHandler(jobs, results))

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
