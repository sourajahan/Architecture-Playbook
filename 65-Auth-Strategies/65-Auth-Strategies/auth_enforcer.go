package security

import (
	"context"
	"fmt"
	"net/http"
)

// AuthzRequest represents the query sent to the Policy Engine.
type AuthzRequest struct {
	Subject string // User ID
	Action  string // READ, WRITE, DELETE
	Object  string // Resource ID
}

// OPAClient simulates the communication with Open Policy Agent.
// This is how we decouple logic from enforcement.
type OPAClient struct {
	Endpoint string
}

func (c *OPAClient) IsAuthorized(ctx context.Context, req AuthzRequest) (bool, error) {
	// In production, this calls a sidecar agent (localhost:8181) via gRPC/HTTP
	fmt.Printf("Querying OPA for: User[%s] wants to %s on %s\n", req.Subject, req.Action, req.Object)
	
	// Mock: Decision is made by centralized policy code, not service code.
	return true, nil 
}

// Middleware: Principal-grade authorization
func (c *OPAClient) AuthorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Extract Identity from validated JWT (previous steps)
		userID := r.Header.Get("X-User-ID") 
		
		// 2. Query Policy Engine
		allowed, err := c.IsAuthorized(r.Context(), AuthzRequest{
			Subject: userID,
			Action:  r.Method,
			Object:  r.URL.Path,
		})
		
		if err != nil || !allowed {
			http.Error(w, "Forbidden: Policy Violation", http.StatusForbidden)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}
