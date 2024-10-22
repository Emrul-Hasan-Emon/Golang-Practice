package main

import (
	"fmt"
	"time"
)

func workerPool(numOfWorkers int, jobs <-chan int, results chan<- int) {
	for worker := 1; worker <= numOfWorkers; worker++ {
		go work(worker, jobs, results)
	}
}

func work(worker int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d picked job %d\n", worker, job)
		results <- process(job)
		fmt.Printf("Worker %d completed job %d\n", worker, job)
	}
}

func process(job int) int {
	// complete the job
	time.Sleep(time.Millisecond * 2)
	return job * 2
}

func main() {
	totalJobs := 12
	numOfWorkers := 5
	jobsQueue := make(chan int, totalJobs)
	results := make(chan int, totalJobs)

	// Start worker pool
	workerPool(numOfWorkers, jobsQueue, results)

	// Send the jobs
	for job := 0; job < totalJobs; job++ {
		jobsQueue <- job
	}

	// Collect Results
	for i := 0; i < totalJobs; i++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
	fmt.Println("Done.........")
}
