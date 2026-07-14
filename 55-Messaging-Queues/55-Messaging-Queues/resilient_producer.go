package main

import (
	"context"
	"fmt"
	"time"
)

// MessageProducer abstracts the broker (Kafka/SQS/Rabbit)
type MessageProducer interface {
	Publish(ctx context.Context, topic string, payload []byte) error
}

// ResilientProducer adds logic for retries and fault tolerance
type ResilientProducer struct {
	broker MessageProducer
}

func (p *ResilientProducer) SendWithRetry(ctx context.Context, topic string, payload []byte) error {
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		err := p.broker.Publish(ctx, topic, payload)
		if err == nil {
			return nil
		}
		
		fmt.Printf("Publish failed, retrying (%d/%d)... \n", i+1, maxRetries)
		// Exponential Backoff
		select {
		case <-time.After(time.Duration(i+1) * 500 * time.Millisecond):
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return fmt.Errorf("exhausted retries")
}
