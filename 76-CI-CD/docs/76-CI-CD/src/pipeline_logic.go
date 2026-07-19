package main

import (
	"fmt"
	"time"
)

// Pipeline represents the automated CI/CD flow
type Pipeline struct {
	ServiceName string
}

// Build compiles the source code and creates artifacts
func (p *Pipeline) Build() error {
	fmt.Printf("[%s] Building artifacts...\n", p.ServiceName)
	time.Sleep(1 * time.Second)
	return nil
}

// Test executes automated validation suites
func (p *Pipeline) Test() error {
	fmt.Printf("[%s] Running integration tests...\n", p.ServiceName)
	time.Sleep(1 * time.Second)
	return nil
}

// Deploy triggers the CD process to the target environment
func (p *Pipeline) Deploy(env string) error {
	fmt.Printf("[%s] Deploying to %s environment...\n", p.ServiceName, env)
	time.Sleep(1 * time.Second)
	return nil
}

func main() {
	// Simple CI/CD pipeline simulation
	service := Pipeline{ServiceName: "Order-Processing-Service"}

	// Execute Pipeline
	if err := service.Build(); err == nil {
		if err := service.Test(); err == nil {
			service.Deploy("Production")
			fmt.Println("Pipeline execution successful.")
		}
	}
}
