package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Pipeline represents a stage in our Reactive flow
type Pipeline struct {
	wg sync.WaitGroup
}

// Source produces data (The Producer)
func (p *Pipeline) Source(ctx context.Context, data []int) <-chan int {
	out := make(chan int, 10) // Buffer size acts as Backpressure mechanism
	go func() {
		defer close(out)
		for _, v := range data {
			select {
			case <-ctx.Done(): return
			case out <- v:
			}
		}
	}()
	return out
}

// Transformer processes data (The Operator)
func (p *Pipeline) Transformer(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int, 10)
	go func() {
		defer close(out)
		for v := range in {
			// Simulate heavy compute/transformation
			out <- v * 2 
		}
	}()
	return out
}

// Sink consumes data (The Subscriber)
func (p *Pipeline) Sink(ctx context.Context, in <-chan int) {
	for v := range in {
		fmt.Printf("Processed event: %d\n", v)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := &Pipeline{}
	
	// Declarative Composition
	source := p.Source(ctx, []int{1, 2, 3, 4, 5})
	transformed := p.Transformer(ctx, source)
	
	p.Sink(ctx, transformed)
}
