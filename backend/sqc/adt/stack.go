package adt

import (
	"errors"
	"sync"
)

// Stack ADT - Abstract Data Type for LIFO operations
// Demonstrates encapsulation and generic data structure design

// Stack interface defines the abstract operations
type Stack[T any] interface {
	Push(item T)
	Pop() (T, error)
	Peek() (T, error)
	IsEmpty() bool
	Size() int
	Clear()
}

// ArrayStack implements Stack using slice
type ArrayStack[T any] struct {
	items []T
	mu    sync.RWMutex // Thread-safe operations
}

// NewStack creates a new stack instance
func NewStack[T any]() Stack[T] {
	return &ArrayStack[T]{
		items: make([]T, 0),
	}
}

// Push adds an item to the top of the stack
func (s *ArrayStack[T]) Push(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, item)
}

// Pop removes and returns the top item
func (s *ArrayStack[T]) Pop() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	var zero T
	if len(s.items) == 0 {
		return zero, errors.New("stack is empty")
	}
	
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

// Peek returns the top item without removing it
func (s *ArrayStack[T]) Peek() (T, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	var zero T
	if len(s.items) == 0 {
		return zero, errors.New("stack is empty")
	}
	
	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty
func (s *ArrayStack[T]) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *ArrayStack[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.items)
}

// Clear removes all items from the stack
func (s *ArrayStack[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = make([]T, 0)
}

