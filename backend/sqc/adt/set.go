package adt

import "sync"

// Set ADT - Abstract Data Type for unique collections
// Demonstrates set theory operations

// Set interface defines set operations
type Set[T comparable] interface {
	Add(item T) bool
	Remove(item T) bool
	Contains(item T) bool
	Size() int
	Clear()
	Items() []T
	Union(other Set[T]) Set[T]
	Intersection(other Set[T]) Set[T]
	Difference(other Set[T]) Set[T]
}

// HashSet implements Set using map
type HashSet[T comparable] struct {
	items map[T]struct{}
	mu    sync.RWMutex
}

// NewSet creates a new set instance
func NewSet[T comparable]() Set[T] {
	return &HashSet[T]{
		items: make(map[T]struct{}),
	}
}

// Add inserts an item into the set
func (s *HashSet[T]) Add(item T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if _, exists := s.items[item]; exists {
		return false
	}
	s.items[item] = struct{}{}
	return true
}

// Remove deletes an item from the set
func (s *HashSet[T]) Remove(item T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if _, exists := s.items[item]; !exists {
		return false
	}
	delete(s.items, item)
	return true
}

// Contains checks if an item exists in the set
func (s *HashSet[T]) Contains(item T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	_, exists := s.items[item]
	return exists
}

// Size returns the number of items
func (s *HashSet[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.items)
}

// Clear removes all items
func (s *HashSet[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = make(map[T]struct{})
}

// Items returns all items as a slice
func (s *HashSet[T]) Items() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	result := make([]T, 0, len(s.items))
	for item := range s.items {
		result = append(result, item)
	}
	return result
}

// Union returns a new set with all items from both sets
func (s *HashSet[T]) Union(other Set[T]) Set[T] {
	result := NewSet[T]()
	
	for _, item := range s.Items() {
		result.Add(item)
	}
	for _, item := range other.Items() {
		result.Add(item)
	}
	
	return result
}

// Intersection returns a new set with common items
func (s *HashSet[T]) Intersection(other Set[T]) Set[T] {
	result := NewSet[T]()
	
	for _, item := range s.Items() {
		if other.Contains(item) {
			result.Add(item)
		}
	}
	
	return result
}

// Difference returns items in s but not in other
func (s *HashSet[T]) Difference(other Set[T]) Set[T] {
	result := NewSet[T]()
	
	for _, item := range s.Items() {
		if !other.Contains(item) {
			result.Add(item)
		}
	}
	
	return result
}

