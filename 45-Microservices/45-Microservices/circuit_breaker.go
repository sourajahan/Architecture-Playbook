package main

import (
	"fmt"
	"sync"
	"time"
)

type State int
const (
	Closed State = iota // Service is healthy
	Open                // Service is failing (Stop traffic)
	HalfOpen            // Testing for recovery
)

// CircuitBreaker manages resilience in a Microservices ecosystem.
type CircuitBreaker struct {
	state      State
	failures   int
	threshold  int
	mu         sync.Mutex
}

// Call wraps a remote service request.
func (cb *CircuitBreaker) Call(serviceFn func() error) error {
	cb.mu.Lock()
	if cb.state == Open {
		cb.mu.Unlock()
		return fmt.Errorf("circuit open: service unavailable")
	}
	cb.mu.Unlock()

	err := serviceFn()

	cb.mu.Lock()
	defer cb.mu.Unlock()
	if err != nil {
		cb.failures++
		if cb.failures >= cb.threshold {
			cb.state = Open
		}
		return err
	}
	cb.failures = 0 // Reset on success
	return nil
}
