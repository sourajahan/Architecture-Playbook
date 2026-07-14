package main

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// NewPrincipalTransport creates a production-grade transport.
// It optimizes connection reuse, TLS negotiation, and resource management.
func NewPrincipalTransport() *http.Transport {
	return &http.Transport{
		// 1. Connection Pooling: Scale idle connections for burst traffic.
		MaxIdleConns:        1000,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     90 * time.Second,

		// 2. Dialing Tuning: Reduce latency for TCP/TLS handshakes.
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,

		// 3. TLS Tuning: Enforce modern TLS 1.3
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS13, 
		},

		// 4. Performance: Enable HTTP/2 multiplexing, force disable for HTTP/1.1
		ForceAttemptHTTP2: true,
	}
}
