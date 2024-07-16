package main

import "fmt"

func main() {
	n := 20

	tasks := make([]Task, n)

	for i := 0; i < n; i++ {
		tasks[i] = Task{ID: i + 1}
	}

	// Create a worker pool
	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5,
	}

	wp.Run()
	fmt.Println("All tasks have been processed!")
}
