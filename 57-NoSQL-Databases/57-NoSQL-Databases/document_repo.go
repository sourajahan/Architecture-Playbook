package main

import "context"

// In NoSQL, we don't Join. We embed.
type UserProfile struct {
    UserID    string      `json:"user_id"`
    Orders    []OrderInfo `json:"orders"` // Denormalization: Embedding orders inside User
}

type OrderInfo struct {
    OrderID string `json:"order_id"`
    Total   float64 `json:"total"`
}

// DocumentRepository handles specific query patterns.
type DocumentRepository interface {
    // We model for the Query, not for the Entity
    GetUserDataWithOrders(ctx context.Context, userID string) (UserProfile, error)
}
