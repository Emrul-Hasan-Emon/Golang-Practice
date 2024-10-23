package main

import (
	"fmt"
	"sync"
)

func fanOut(
	numOfWorkers int,
	subTasksQueue chan int,
	results chan<- int,
	subTasks []int,
	wg *sync.WaitGroup,
) {
	defer close(subTasksQueue)

	// Send the tasks into the channel (Fan-Out)
	for _, task := range subTasks {
		subTasksQueue <- task
	}

	// Start the workers
	for w := 1; w <= numOfWorkers; w++ {
		wg.Add(1)
		go worker(w, subTasksQueue, results, wg)
	}
}

func worker(w int, subTasksQueue <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Process the tasks from the queue
	for subTask := range subTasksQueue {
		fmt.Printf("Worker %d started finding square of %d subTask\n", w, subTask)
		results <- process(subTask)
		fmt.Printf("Worker %d completed finding square of %d subTask\n", w, subTask)
	}
}

func process(num int) int {
	return num * num
}

// Collect the results (Fan-In) and then find the summation
func fanIn(results <-chan int, numOfSubTasks int) int {
	square := []int{}
	// Collect the results
	for i := 0; i < numOfSubTasks; i++ {
		result := <-results
		square = append(square, result)
	}

	// Find the summation
	sum := 0
	for _, sq := range square {
		sum += sq
	}
	return sum
}

func main() {
	subTasks := []int{1, 2, 3, 4, 5}
	numOfSubTasks := len(subTasks)
	numOfWorkers := 2

	subTasksQueue := make(chan int, numOfSubTasks)
	results := make(chan int, numOfSubTasks)

	var wg sync.WaitGroup

	// Start the fan-out in a separate goroutine
	go fanOut(numOfWorkers, subTasksQueue, results, subTasks, &wg)

	// We need to wait until all the sub tasks are completed
	wg.Wait()

	// Fan-in to collect the results and find the summation
	sum := fanIn(results, numOfSubTasks)
	close(results)

	fmt.Printf("Total sum of squares: %d\n", sum)
}
