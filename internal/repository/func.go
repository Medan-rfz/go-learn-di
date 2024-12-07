package repository

import (
	"context"
	"log"
)

// Func ...
func (r *Repo) Func(ctx context.Context) (err error) {
	// TODO Some logic
	log.Printf("Repo: was calling Func() with connStr '%s'", r.conn)
	return
}
