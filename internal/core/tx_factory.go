package core

import "context"

type IQueryable interface {
	// Commits the current transaction (performs rollback on
	// failure to commit) and begins a new internal transaction.
	Save() error
}

type IDbTransaction interface {
	IQueryable

	Commit() error
	Rollback()
}

type IDbTransactionFactory interface {
	NewTx(ctx context.Context) (IDbTransaction, error)
	WithTx(ctx context.Context, fn func(IQueryable) error) error
}
