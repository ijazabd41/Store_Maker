package concurrency

import (
	"context"
	"sync"
)

// WorkerPool - Demonstrates concurrency pattern
// Implements a worker pool for parallel task processing

// Task represents a unit of work
type Task func() error

// Result holds the result of a task execution
type Result struct {
	Index int
	Error error
	Data  interface{}
}

// WorkerPool manages concurrent workers
type WorkerPool struct {
	workers    int
	taskQueue  chan Task
	resultChan chan Result
	wg         sync.WaitGroup
	ctx        context.Context
	cancel     context.CancelFunc
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(workers int, queueSize int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	
	return &WorkerPool{
		workers:    workers,
		taskQueue:  make(chan Task, queueSize),
		resultChan: make(chan Result, queueSize),
		ctx:        ctx,
		cancel:     cancel,
	}
}

// Start initializes and starts all workers
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// worker is the worker goroutine
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	
	for {
		select {
		case task, ok := <-wp.taskQueue:
			if !ok {
				return // Channel closed, worker exits
			}
			
			// Execute task
			err := task()
			
			// Send result (if needed)
			select {
			case wp.resultChan <- Result{Index: id, Error: err}:
			case <-wp.ctx.Done():
				return
			}
			
		case <-wp.ctx.Done():
			return // Context canceled, worker exits
		}
	}
}

// Submit adds a task to the queue
func (wp *WorkerPool) Submit(task Task) bool {
	select {
	case wp.taskQueue <- task:
		return true
	case <-wp.ctx.Done():
		return false
	default:
		return false // Queue full
	}
}

// Stop gracefully stops the worker pool
func (wp *WorkerPool) Stop() {
	close(wp.taskQueue)
	wp.wg.Wait()
	close(wp.resultChan)
}

// Results returns the result channel
func (wp *WorkerPool) Results() <-chan Result {
	return wp.resultChan
}

// BatchProcess processes items in parallel using the worker pool
func BatchProcess[T any](items []T, workers int, processFunc func(T) error) []error {
	pool := NewWorkerPool(workers, len(items))
	pool.Start()
	
	errors := make([]error, len(items))
	var errorsMu sync.Mutex
	
	// Submit all tasks
	for idx, item := range items {
		index := idx
		currentItem := item
		
		task := func() error {
			err := processFunc(currentItem)
			if err != nil {
				errorsMu.Lock()
				errors[index] = err
				errorsMu.Unlock()
			}
			return err
		}
		
		pool.Submit(task)
	}
	
	// Wait for completion
	pool.Stop()
	
	return errors
}

