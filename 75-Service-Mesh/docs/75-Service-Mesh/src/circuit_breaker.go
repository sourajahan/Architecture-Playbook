package main

import (
	"errors"
	"fmt"
	"time"
)

// CircuitState defines the operational mode of the protected service.
type CircuitState int

const (
	StateClosed CircuitState = iota // Flow: Normal operation
	StateOpen                       // Flow: Blocked to prevent cascading failure
	StateHalfOpen                   // Flow: Probing for service recovery
)

// CircuitBreaker patterns encapsulate failure detection logic outside of business code.
type CircuitBreaker struct {
	state             CircuitState
	failureCount      int
	failureThreshold  int
	resetTimeout      time.Duration
	lastFailureTime   time.Time
}

func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:            StateClosed,
		failureThreshold: threshold,
		resetTimeout:     timeout,
	}
}

// Execute enforces the safety policy by wrapping external service calls.
func (cb *CircuitBreaker) Execute(requestFunc func() error) error {
	if cb.state == StateOpen {
		if time.Since(cb.lastFailureTime) > cb.resetTimeout {
			cb.state = StateHalfOpen
		} else {
			return errors.New("circuit breaker is OPEN: Request blocked")
		}
	}

	err := requestFunc()
	if err != nil {
		cb.recordFailure()
		return err
	}

	cb.reset()
	return nil
}

func (cb *CircuitBreaker) recordFailure() {
	cb.failureCount++
	cb.lastFailureTime = time.Now()
	if cb.failureCount >= cb.failureThreshold {
		cb.state = StateOpen
	}
}

func (cb *CircuitBreaker) reset() {
	cb.state = StateClosed
	cb.failureCount = 0
}
