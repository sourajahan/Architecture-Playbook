package serverless

import (
	"context"
	"fmt"
	"log"
)

// ServerlessHandler enforces structural integrity.
type ServerlessHandler struct {
	QueueProducer QueueClient // Used to send failed events to DLQ
}

// HandleEvent processes an incoming event with safety checks.
func (h *ServerlessHandler) HandleEvent(ctx context.Context, event RawEvent) error {
	// 1. Validation (Defensive programming is mandatory)
	if !event.IsValid() {
		return fmt.Errorf("invalid event payload")
	}

	// 2. Execution (Business logic)
	err := h.process(ctx, event)
	if err != nil {
		// 3. Resilience: If processing fails, push to DLQ for manual inspection
		log.Printf("Processing failed, routing to DLQ: %v", err)
		h.QueueProducer.SendToDLQ(event) 
		return nil // Return nil so the event source (SQS/Kafka) considers it "consumed"
	}

	return nil
}

func (h *ServerlessHandler) process(ctx context.Context, event RawEvent) error {
	// Logic goes here
	return nil
}
