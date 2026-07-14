package main

import (
	"context"
	"fmt"
	"time"
)

// Client handles the communication logic.
type Client struct {
	BaseURL string
}

// FetchDataWithRetry is a Principal-level pattern to handle transient network failures.
func (c *Client) FetchDataWithRetry(ctx context.Context, endpoint string, retries int) (string, error) {
	var lastErr error
	for i := 0; i < retries; i++ {
		// Attempting to simulate a network call
		data, err := c.performRequest(ctx, endpoint)
		if err == nil {
			return data, nil
		}
		
		lastErr = err
		// Exponential Backoff: Wait longer with each failure
		backoff := time.Duration(1<<i) * 100 * time.Millisecond
		fmt.Printf("Attempt %d failed, retrying in %v...\n", i+1, backoff)
		
		select {
		case <-time.After(backoff):
		case <-ctx.Done(): // Respect context cancellation
			return "", ctx.Err()
		}
	}
	return "", fmt.Errorf("failed after %d retries: %w", retries, lastErr)
}

func (c *Client) performRequest(ctx context.Context, endpoint string) (string, error) {
	// Simulate unstable network
	return "Success", nil 
}
