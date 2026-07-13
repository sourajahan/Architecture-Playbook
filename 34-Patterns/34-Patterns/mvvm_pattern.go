package main

import (
	"fmt"
	"sync"
)

// Observer is the interface for UI components
type Observer interface {
	Update(state interface{})
}

// Store holds the State (The "ViewModel" core)
type Store struct {
	state     interface{}
	observers []Observer
	mu        sync.RWMutex
}

func (s *Store) Subscribe(o Observer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.observers = append(s.observers, o)
}

// Dispatch transforms state and notifies observers (Reactive Binding)
func (s *Store) Dispatch(newState interface{}) {
	s.mu.Lock()
	s.state = newState
	s.mu.Unlock()

	for _, o := range s.observers {
		o.Update(newState)
	}
}

// View implementation
type DashboardView struct {
	Name string
}

func (d *DashboardView) Update(state interface{}) {
	fmt.Printf("[View: %s] Rendering new state: %v\n", d.Name, state)
}

func main() {
	// Centralized Source of Truth
	store := &Store{}

	// Attaching Components (Decoupled)
	view1 := &DashboardView{Name: "Header"}
	view2 := &DashboardView{Name: "Sidebar"}

	store.Subscribe(view1)
	store.Subscribe(view2)

	// State Mutation triggers automatic UI update (Reactivity)
	store.Dispatch("UserAuthenticated: Admin")
}
