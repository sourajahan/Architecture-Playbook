package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Task represents the unit of work
type Task struct {
	ID      int
	Payload string
}

// Worker processes tasks from the job channel
func worker(ctx context.Context, id int, jobs <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: shutting down\n", id)
				return
			}
			fmt.Printf("Worker %d: processing task %d (%s)\n", id, task.ID, task.Payload)
			time.Sleep(500 * time.Millisecond) // Simulate heavy work
		case <-ctx.Done():
			fmt.Printf("Worker %d: context cancelled, exiting\n", id)
			return
		}
	}
}

func main() {
	const workerCount = 3
	jobs := make(chan Task, 10)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Spawn workers
	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, &wg)
	}

	// Dispatch tasks
	for j := 1; j <= 5; j++ {
		jobs <- Task{ID: j, Payload: "Data"}
	}
	
	close(jobs) // Signal workers that no more jobs will be sent
	wg.Wait()   // Wait for all workers to finish
	fmt.Println("All tasks processed. System shutting down.")
}
