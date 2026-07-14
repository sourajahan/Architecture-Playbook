package main

import (
	"fmt"
	"sync"
)

// Service definition for our SOA Ecosystem
type Service struct {
	Name     string
	Endpoint string
}

// Registry acts as the "Yellow Pages" for Enterprise Services.
// This allows reuse across different teams.
type Registry struct {
	mu       sync.RWMutex
	services map[string]Service
}

func (r *Registry) Register(s Service) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.services[s.Name] = s
	fmt.Printf("Service [%s] registered at %s\n", s.Name, s.Endpoint)
}

func (r *Registry) Discover(name string) (Service, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	s, ok := r.services[name]
	return s, ok
}
