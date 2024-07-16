package main

import (
	"fmt"
	"sync"
	"time"
)

// Task definition
type Task interface {
	Process()
}

// Email task definition
type EmailTask struct {
	Email       string
	Subject     string
	MessageBody string
}

// Image processing task
type ImageProcessingTask struct {
	ImageUrl string
}

// Way to process the Email tasks
func (t *EmailTask) Process() {
	fmt.Printf("Sending email to %s\n", t.Email)
	// Simulate a time consuming process
	time.Sleep(2 * time.Second)
}

// Way to process the Image tasks
func (t *ImageProcessingTask) Process() {
	fmt.Printf("Processing the image %s\n", t.ImageUrl)
	// Simulate a time consuming process
	time.Sleep(4 * time.Second)
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
