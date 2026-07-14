package main

import (
	"context"
	"errors"
)

// 1. DDD: Aggregate Root
// The Order entity encapsulates state and enforces business rules (Invariants).
type Order struct {
	ID     string
	Status string
	Total  float64
}

func (o *Order) ProcessPayment(amount float64) error {
	if o.Status == "PAID" {
		return errors.New("domain error: order already paid")
	}
	o.Status = "PAID"
	return nil
}

// 2. SOLID: Dependency Inversion
// The Service depends on an interface, not an implementation (SQL/NoSQL).
type OrderRepository interface {
	Save(ctx context.Context, order *Order) error
	GetByID(ctx context.Context, id string) (*Order, error)
}

// 3. TDD-Friendly Service
// Business Logic is isolated and easily testable by mocking the repository.
type OrderService struct {
	repo OrderRepository
}

func (s *OrderService) CompleteOrder(ctx context.Context, id string, payment float64) error {
	order, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := order.ProcessPayment(payment); err != nil {
		return err
	}

	return s.repo.Save(ctx, order)
}
