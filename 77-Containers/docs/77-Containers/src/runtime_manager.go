// Package runtime handles container lifecycle abstractions.
package main

import (
	"fmt"
	"time"
)

// Container represents an isolated workload instance.
type Container struct {
	ID        string
	Resources map[string]string // CPU, Memory constraints
	Status    string
}

// NewContainer initializes a container with defined resource constraints.
func NewContainer(id string, cpu string) *Container {
	return &Container{
		ID:     id,
		Resources: map[string]string{"cpu": cpu},
		Status: "CREATED",
	}
}

// Run executes the containerized process.
func (c *Container) Run() {
	c.Status = "RUNNING"
	fmt.Printf("[Container %s] Process started with resource profile: %v\n", c.ID, c.Resources)
}

func main() {
	ct := NewContainer("prod-web-01", "2000m")
	ct.Run()
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("[System] Container %s is healthy.\n", ct.ID)
}
