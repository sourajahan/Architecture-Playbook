package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Task represents a unit of work in our distributed node.
type Task struct {
	ID int
}

// Worker processes tasks. In a real system, this talks to a queue (Kafka/RabbitMQ).
func Worker(ctx context.Context, id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		select {
		case <-ctx.Done():
			return // Graceful shutdown
		default:
			// Process with potential failure simulation
			fmt.Printf("Worker %d processing task %d\n", id, task.ID)
			time.Sleep(100 * time.Millisecond) 
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	tasks := make(chan Task, 10)
	var wg sync.WaitGroup

	// Scaling: Spin up workers dynamically
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go Worker(ctx, i, tasks, &wg)
	}

	// Producer: Injecting load
	for i := 1; i <= 5; i++ {
		tasks <- Task{ID: i}
	}
	close(tasks)
	wg.Wait()
}
