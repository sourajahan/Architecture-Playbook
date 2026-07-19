// Package governance provides tools for validating project compliance with industry standards.
package main

import (
	"fmt"
)

// ComplianceStatus represents the adherence level of a project.
type ComplianceStatus string

const (
	Compliant    ComplianceStatus = "COMPLIANT"
	NonCompliant ComplianceStatus = "NON-COMPLIANT"
)

// ProjectGovernance represents the state of a project against framework requirements.
type ProjectGovernance struct {
	ProjectName       string
	HasRiskAssessment bool // PMI requirement
	HasSLADefined     bool // ITIL requirement
	IsIterative       bool // RUP requirement
}

// ValidateFramework ensures the project meets the minimum standards of a specific methodology.
func (pg *ProjectGovernance) ValidateFramework(framework string) ComplianceStatus {
	fmt.Printf("[Governance Audit] Checking framework: %s for Project: %s\n", framework, pg.ProjectName)
	
	switch framework {
	case "ITIL":
		if pg.HasSLADefined { return Compliant }
	case "PMI":
		if pg.HasRiskAssessment { return Compliant }
	case "RUP":
		if pg.IsIterative { return Compliant }
	}
	return NonCompliant
}

func main() {
	// Audit a project instance
	project := ProjectGovernance{
		ProjectName:       "Enterprise-Cloud-Migration",
		HasRiskAssessment: true,
		HasSLADefined:     true,
		IsIterative:       true,
	}

	// Validate against multiple standards
	standards := []string{"ITIL", "PMI", "RUP"}
	
	for _, std := range standards {
		status := project.ValidateFramework(std)
		fmt.Printf("Audit Result: [%s] -> %s\n", std, status)
	}
}
