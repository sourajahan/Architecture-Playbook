// Package agile provides utilities for calculating sprint metrics and team performance.
package main

import (
	"fmt"
)

// SprintData represents the outcomes of an iteration.
type SprintData struct {
	SprintName        string
	CompletedStoryPoints int
}

// VelocityCalculator manages team performance analytics.
type VelocityCalculator struct {
	History []SprintData
}

// CalculateAverageVelocity computes the rolling average velocity of the team.
func (vc *VelocityCalculator) CalculateAverageVelocity() float64 {
	if len(vc.History) == 0 {
		return 0
	}
	
	total := 0
	for _, sprint := range vc.History {
		total += sprint.CompletedStoryPoints
	}
	return float64(total) / float64(len(vc.History))
}

func main() {
	// Initialize calculator with past sprint data
	calc := VelocityCalculator{
		History: []SprintData{
			{"Sprint-1", 20},
			{"Sprint-2", 25},
			{"Sprint-3", 22},
		},
	}

	avg := calc.CalculateAverageVelocity()
	fmt.Printf("[Agile Metrics] Team Average Velocity: %.2f Story Points\n", avg)
	fmt.Println("[Agile Metrics] Use this velocity to forecast future capacity.")
}
