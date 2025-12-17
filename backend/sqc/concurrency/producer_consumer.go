package concurrency

import (
	"context"
	"sync"
)

// ProducerConsumer - Classic concurrency pattern
// Demonstrates bounded buffer problem solution

// Buffer represents a thread-safe bounded buffer
type Buffer[T any] struct {
	data     []T
	capacity int
	size     int
	head     int
	tail     int
	mu       sync.Mutex
	notEmpty *sync.Cond
	notFull  *sync.Cond
}

// NewBuffer creates a new bounded buffer
func NewBuffer[T any](capacity int) *Buffer[T] {
	b := &Buffer[T]{
		data:     make([]T, capacity),
		capacity: capacity,
		size:     0,
		head:     0,
		tail:     0,
	}
	b.notEmpty = sync.NewCond(&b.mu)
	b.notFull = sync.NewCond(&b.mu)
	return b
}

// Produce adds an item to the buffer (blocks if full)
func (b *Buffer[T]) Produce(item T) {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	// Wait while buffer is full
	for b.size == b.capacity {
		b.notFull.Wait()
	}
	
	// Add item
	b.data[b.tail] = item
	b.tail = (b.tail + 1) % b.capacity
	b.size++
	
	// Signal that buffer is not empty
	b.notEmpty.Signal()
}

// Consume removes and returns an item (blocks if empty)
func (b *Buffer[T]) Consume() T {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	// Wait while buffer is empty
	for b.size == 0 {
		b.notEmpty.Wait()
	}
	
	// Remove item
	item := b.data[b.head]
	b.head = (b.head + 1) % b.capacity
	b.size--
	
	// Signal that buffer is not full
	b.notFull.Signal()
	
	return item
}

// TryProduce attempts to add item without blocking
func (b *Buffer[T]) TryProduce(item T) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	if b.size == b.capacity {
		return false
	}
	
	b.data[b.tail] = item
	b.tail = (b.tail + 1) % b.capacity
	b.size++
	b.notEmpty.Signal()
	
	return true
}

// TryConsume attempts to remove item without blocking
func (b *Buffer[T]) TryConsume() (T, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	var zero T
	if b.size == 0 {
		return zero, false
	}
	
	item := b.data[b.head]
	b.head = (b.head + 1) % b.capacity
	b.size--
	b.notFull.Signal()
	
	return item, true
}

// Size returns current buffer size
func (b *Buffer[T]) Size() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.size
}

// ProducerConsumerSystem manages multiple producers and consumers
type ProducerConsumerSystem[T any] struct {
	buffer    *Buffer[T]
	ctx       context.Context
	cancel    context.CancelFunc
	producers int
	consumers int
	wg        sync.WaitGroup
}

// NewProducerConsumerSystem creates a new system
func NewProducerConsumerSystem[T any](bufferSize int) *ProducerConsumerSystem[T] {
	ctx, cancel := context.WithCancel(context.Background())
	
	return &ProducerConsumerSystem[T]{
		buffer: NewBuffer[T](bufferSize),
		ctx:    ctx,
		cancel: cancel,
	}
}

// StartProducer starts a producer goroutine
func (pcs *ProducerConsumerSystem[T]) StartProducer(produceFunc func() T) {
	pcs.wg.Add(1)
	pcs.producers++
	
	go func() {
		defer pcs.wg.Done()
		
		for {
			select {
			case <-pcs.ctx.Done():
				return
			default:
				item := produceFunc()
				pcs.buffer.Produce(item)
			}
		}
	}()
}

// StartConsumer starts a consumer goroutine
func (pcs *ProducerConsumerSystem[T]) StartConsumer(consumeFunc func(T)) {
	pcs.wg.Add(1)
	pcs.consumers++
	
	go func() {
		defer pcs.wg.Done()
		
		for {
			select {
			case <-pcs.ctx.Done():
				return
			default:
				item := pcs.buffer.Consume()
				consumeFunc(item)
			}
		}
	}()
}

// Stop gracefully stops the system
func (pcs *ProducerConsumerSystem[T]) Stop() {
	pcs.cancel()
	pcs.wg.Wait()
}

// GetBufferSize returns current buffer size
func (pcs *ProducerConsumerSystem[T]) GetBufferSize() int {
	return pcs.buffer.Size()
}

