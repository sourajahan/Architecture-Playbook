package main

// --- LAYER 3: Repository (Infrastructure) ---
// Implementation detail (e.g., PostgreSQL)
type PostgresRepo struct{}

func (p *PostgresRepo) Save(data string) { /* SQL Logic */ }

// --- LAYER 2: Service (Domain Logic) ---
// This layer is agnostic of the DB. It uses the Repository Interface.
type Repository interface {
	Save(data string)
}

type OrderService struct {
	repo Repository // Dependency Inversion in action
}

func (s *OrderService) CreateOrder(data string) {
	// Business rules here...
	s.repo.Save(data)
}

// --- LAYER 1: Presentation (Transport) ---
// HTTP Handler receives the request and delegates to the Service.
type OrderHandler struct {
	service *OrderService
}

func (h *OrderHandler) HandleHTTP(data string) {
	h.service.CreateOrder(data)
}
