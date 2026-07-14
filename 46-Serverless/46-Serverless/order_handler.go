package main

import (
	"context"
	"fmt"
)

// Event represents the trigger payload (e.g., from AWS EventBridge/SQS)
type OrderEvent struct {
	OrderID string `json:"order_id"`
	Amount  float64 `json:"amount"`
}

// OrderHandler is an ephemeral execution unit.
// It has no memory of previous executions (Stateless).
func OrderHandler(ctx context.Context, event OrderEvent) error {
	// 1. Validation (Business Logic)
	if event.Amount <= 0 {
		return fmt.Errorf("invalid order amount")
	}

	// 2. Interaction with external state (e.g., DynamoDB/Postgres)
	// We MUST persist state here to survive the function lifecycle.
	fmt.Printf("Processing order %s for $%.2f\n", event.OrderID, event.Amount)

	// 3. Trigger downstream events (Choreography)
	// We don't call other services directly; we emit events.
	return nil 
}
