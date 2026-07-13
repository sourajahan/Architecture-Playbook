package main

import (
	"context"
	"database/sql"
	"fmt"
)

// CommandHandler manages the Domain logic and Atomicity.
type CommandHandler struct {
	db *sql.DB
}

// ProcessCommand executes an atomic state update and outbox insertion.
func (ch *CommandHandler) ProcessCommand(ctx context.Context, aggregateID, payload string) error {
	// Start an atomic transaction.
	tx, err := ch.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 1. Update the Domain Aggregate.
	_, err = tx.ExecContext(ctx, "UPDATE aggregates SET data = ? WHERE id = ?", payload, aggregateID)
	if err != nil {
		return err
	}

	// 2. Insert into Outbox table within the same transaction.
	// This guarantees that if the update fails, the event is not created.
	_, err = tx.ExecContext(ctx, "INSERT INTO outbox (aggregate_id, event_type, payload) VALUES (?, ?, ?)", 
		aggregateID, "AggregateUpdated", payload)
	if err != nil {
		return err
	}

	// 3. Atomically commit both.
	return tx.Commit()
}

// RelayService simulates the polling mechanism for the Outbox.
func (ch *CommandHandler) RelayService(ctx context.Context) {
	// 1. SELECT * FROM outbox WHERE processed = false LIMIT 10
	// 2. Publish to Message Broker (e.g., Kafka).
	// 3. Mark as processed or delete from outbox.
	fmt.Println("Relay: Guaranteed delivery of events to downstream subscribers.")
}
