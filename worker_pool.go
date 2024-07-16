package main

import (
	"fmt"
	"sync"
	"time"
)

// Task definition
type Task struct {
	ID int
}

// Way to process the tasks
func (t *Task) Process() {
	fmt.Printf("Processing task %d\n", t.ID)
	// Simulate a time consuming process
	time.Sleep(2 * time.Second)
}

// Worker pool definition
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	tasksChan   chan Task
	wg          sync.WaitGroup
}

// Functions to execute the worker pool
func (wp *WorkerPool) worker() {
	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	n := len(wp.Tasks)
	wp.tasksChan = make(chan Task, n)

	// Start workers
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	// Send tasks to the tasks channel
	wp.wg.Add(n)
	for _, task := range wp.Tasks {
		wp.tasksChan <- task
	}

	close(wp.tasksChan)

	wp.wg.Wait()
}
