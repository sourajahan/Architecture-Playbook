// Package workflow orchestrates document-centric processes.
package main

import (
	"fmt"
	"time"
)

// Document represents an ECM entity with metadata.
type Document struct {
	ID       string
	Status   string // e.g., "Draft", "Approved", "Archived"
	Owner    string
}

// WorkflowEngine manages the BPM state machine.
type WorkflowEngine struct {
	EngineName string
}

// ProcessDocument simulates a workflow transition (e.g., Approval Cycle).
func (we *WorkflowEngine) ProcessDocument(doc *Document, action string) error {
	fmt.Printf("[%s] Orchestrating workflow for DocID: %s\n", we.EngineName, doc.ID)
	
	switch action {
	case "SUBMIT":
		doc.Status = "Pending-Approval"
	case "APPROVE":
		doc.Status = "Archived"
	default:
		return fmt.Errorf("invalid action: %s", action)
	}

	fmt.Printf("[BPM] Process completed. Document %s transitioned to %s\n", doc.ID, doc.Status)
	return nil
}

func main() {
	// Initialize IBM-style BPM Engine
	bpm := WorkflowEngine{EngineName: "IBM-BAW-Instance-01"}
	
	// Create an EMC-style Document
	doc := &Document{ID: "DOC-9982", Status: "Draft", Owner: "Finance_Dept"}

	// Execute Orchestration
	err := bpm.ProcessDocument(doc, "SUBMIT")
	if err != nil {
		fmt.Println("Error:", err)
	}
	
	time.Sleep(500 * time.Millisecond)
	bpm.ProcessDocument(doc, "APPROVE")
}
