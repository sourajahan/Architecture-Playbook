package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

// DiagnosticResult captures the health status across different layers.
type DiagnosticResult struct {
	Layer4Status bool    // TCP Connection
	Layer7Status bool    // HTTP/gRPC Response
	LatencyL4    time.Duration
	LatencyL7    time.Duration
}

// NetworkDiagnosticEngine analyzes the path to the downstream service.
type NetworkDiagnosticEngine struct {
	TargetAddr string
	TargetURL  string
}

// RunDiagnostics performs a deep check through the OSI layers.
func (e *NetworkDiagnosticEngine) RunDiagnostics(ctx context.Context) (DiagnosticResult, error) {
	result := DiagnosticResult{}

	// --- LAYER 4 DIAGNOSTICS (TCP/Transport Layer) ---
	startL4 := time.Now()
	conn, err := net.DialTimeout("tcp", e.TargetAddr, 5*time.Second)
	if err != nil {
		return result, fmt.Errorf("Layer 4 Failure (TCP Connection): %w", err)
	}
	defer conn.Close()
	result.Layer4Status = true
	result.LatencyL4 = time.Since(startL4)

	// --- LAYER 7 DIAGNOSTICS (App Layer) ---
	startL7 := time.Now()
	req, _ := http.NewRequestWithContext(ctx, "GET", e.TargetURL, nil)
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	
	if err != nil {
		return result, fmt.Errorf("Layer 7 Failure (Application): %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		result.Layer7Status = true
	}
	result.LatencyL7 = time.Since(startL7)

	return result, nil
}
