package main

import (
	"context"
	"database/sql"
)

// DBRouter implements Read/Write splitting.
type DBRouter struct {
	Master *sql.DB // Write node
	Replica *sql.DB // Read node
}

// Execute performs a write operation (e.g., INSERT, UPDATE)
func (r *DBRouter) Execute(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	// Always route writes to Master
	return r.Master.ExecContext(ctx, query, args...)
}

// Query performs a read operation
func (r *DBRouter) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	// Route reads to Replica for scale
	return r.Replica.QueryContext(ctx, query, args...)
}
