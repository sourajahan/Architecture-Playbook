package main

import (
	"context"
	"fmt"
	"sync"
)

// Command represents the message payload (DTO).
type Command struct {
	Action string
	Data   any
}

// Actor defines the behavioral contract.
type Actor interface {
	Process(ctx context.Context, cmd Command) error
}

// WorkerActor implements the Actor interface with isolated state.
type WorkerActor struct {
	id      int
	mailbox chan Command
}

func (a *WorkerActor) Process(ctx context.Context, cmd Command) error {
	// Business logic is encapsulated here.
	fmt.Printf("[Actor %d] Executing: %s\n", a.id, cmd.Action)
	return nil
}

// Start manages the actor's lifecycle loop.
func (a *WorkerActor) Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Actor %d shutting down.\n", a.id)
			return
		case cmd := <-a.mailbox:
			_ = a.Process(ctx, cmd)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	actor := &WorkerActor{id: 42, mailbox: make(chan Command, 10)}
	
	var wg sync.WaitGroup
	wg.Add(1)
	go actor.Start(ctx, &wg)

	// Sending messages asynchronously (Non-blocking)
	actor.mailbox <- Command{Action: "COMPUTE_HASH", Data: 12345}
	
	cancel() // Trigger graceful shutdown
	wg.Wait()
}
