package adt

import (
	"errors"
	"sync"
)

// Queue ADT - Abstract Data Type for FIFO operations
// Used for task scheduling and batch processing

// Queue interface defines abstract queue operations
type Queue[T any] interface {
	Enqueue(item T)
	Dequeue() (T, error)
	Front() (T, error)
	IsEmpty() bool
	Size() int
	Clear()
}

// LinkedQueue implements Queue using linked list
type LinkedQueue[T any] struct {
	head *node[T]
	tail *node[T]
	size int
	mu   sync.RWMutex
}

type node[T any] struct {
	data T
	next *node[T]
}

// NewQueue creates a new queue instance
func NewQueue[T any]() Queue[T] {
	return &LinkedQueue[T]{}
}

// Enqueue adds an item to the rear of the queue
func (q *LinkedQueue[T]) Enqueue(item T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	
	newNode := &node[T]{data: item}
	
	if q.tail == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.size++
}

// Dequeue removes and returns the front item
func (q *LinkedQueue[T]) Dequeue() (T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	
	var zero T
	if q.head == nil {
		return zero, errors.New("queue is empty")
	}
	
	item := q.head.data
	q.head = q.head.next
	
	if q.head == nil {
		q.tail = nil
	}
	
	q.size--
	return item, nil
}

// Front returns the front item without removing it
func (q *LinkedQueue[T]) Front() (T, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	
	var zero T
	if q.head == nil {
		return zero, errors.New("queue is empty")
	}
	
	return q.head.data, nil
}

// IsEmpty checks if the queue is empty
func (q *LinkedQueue[T]) IsEmpty() bool {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return q.size == 0
}

// Size returns the number of items in the queue
func (q *LinkedQueue[T]) Size() int {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return q.size
}

// Clear removes all items from the queue
func (q *LinkedQueue[T]) Clear() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.head = nil
	q.tail = nil
	q.size = 0
}

