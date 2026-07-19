// Package patterns implements architectural design patterns for cloud environments.
package main

import (
	"fmt"
	"sync"
	"time"
)

// Bulkhead manages resource isolation by limiting concurrent access to specific services.
type Bulkhead struct {
	semaphore chan struct{}
}

// NewBulkhead initializes a bulkhead with a fixed capacity of concurrent slots.
func NewBulkhead(capacity int) *Bulkhead {
	return &Bulkhead{
		semaphore: make(chan struct{}, capacity),
	}
}

// Execute wraps a service call, ensuring it does not exceed the allotted capacity.
func (b *Bulkhead) Execute(taskName string, work func()) bool {
	select {
	case b.semaphore <- struct{}{}:
		defer func() { <-b.semaphore }()
		work()
		return true
	default:
		// Capacity reached, reject the request (Fail Fast)
		fmt.Printf("[Bulkhead] Capacity full for service: %s. Rejecting request.\n", taskName)
		return false
	}
}

func main() {
	// Isolate resources for specific service (e.g., only 2 concurrent requests allowed)
	serviceBulkhead := NewBulkhead(2)
	var wg sync.WaitGroup

	// Simulate 5 incoming requests
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			success := serviceBulkhead.Execute(fmt.Sprintf("Request-%d", id), func() {
				fmt.Printf("[Service] Processing Request-%d\n", id)
				time.Sleep(1 * time.Second)
			})
			if !success {
				fmt.Printf("[Client] Request-%d failed due to bulkhead pressure\n", id)
			}
		}(i)
	}
	wg.Wait()
}
