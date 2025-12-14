package concurrency

import (
	"context"
)

// Pipeline - Implements the pipeline concurrency pattern
// Data flows through multiple stages, each processed concurrently

// Stage represents a processing stage in the pipeline
type Stage[T any] func(context.Context, <-chan T) <-chan T

// Pipeline connects multiple stages
type Pipeline[T any] struct {
	ctx    context.Context
	cancel context.CancelFunc
	stages []Stage[T]
}

// NewPipeline creates a new pipeline
func NewPipeline[T any](ctx context.Context, stages ...Stage[T]) *Pipeline[T] {
	pipelineCtx, cancel := context.WithCancel(ctx)
	
	return &Pipeline[T]{
		ctx:    pipelineCtx,
		cancel: cancel,
		stages: stages,
	}
}

// Execute runs the pipeline
func (p *Pipeline[T]) Execute(input <-chan T) <-chan T {
	out := input
	
	// Chain all stages
	for _, stage := range p.stages {
		out = stage(p.ctx, out)
	}
	
	return out
}

// Stop cancels the pipeline
func (p *Pipeline[T]) Stop() {
	p.cancel()
}

// Generator creates a channel and fills it with data
func Generator[T any](ctx context.Context, items ...T) <-chan T {
	out := make(chan T)
	
	go func() {
		defer close(out)
		for _, item := range items {
			select {
			case out <- item:
			case <-ctx.Done():
				return
			}
		}
	}()
	
	return out
}

// Filter creates a filter stage
func Filter[T any](predicate func(T) bool) Stage[T] {
	return func(ctx context.Context, in <-chan T) <-chan T {
		out := make(chan T)
		
		go func() {
			defer close(out)
			for item := range in {
				if predicate(item) {
					select {
					case out <- item:
					case <-ctx.Done():
						return
					}
				}
			}
		}()
		
		return out
	}
}

// Map creates a map/transform stage
func Map[T any, R any](transform func(T) R) func(context.Context, <-chan T) <-chan R {
	return func(ctx context.Context, in <-chan T) <-chan R {
		out := make(chan R)
		
		go func() {
			defer close(out)
			for item := range in {
				transformed := transform(item)
				select {
				case out <- transformed:
				case <-ctx.Done():
					return
				}
			}
		}()
		
		return out
	}
}

// FanOut duplicates input to multiple outputs
func FanOut[T any](ctx context.Context, in <-chan T, n int) []<-chan T {
	outputs := make([]<-chan T, n)
	
	for i := 0; i < n; i++ {
		out := make(chan T)
		outputs[i] = out
		
		go func(ch chan T) {
			defer close(ch)
			for item := range in {
				select {
				case ch <- item:
				case <-ctx.Done():
					return
				}
			}
		}(out)
	}
	
	return outputs
}

// FanIn merges multiple inputs into one output
func FanIn[T any](ctx context.Context, inputs ...<-chan T) <-chan T {
	out := make(chan T)
	
	go func() {
		defer close(out)
		
		for _, input := range inputs {
			for item := range input {
				select {
				case out <- item:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	
	return out
}

