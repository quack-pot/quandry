package core

import "context"

type IQueryable interface {
	/** TODO: Add identifying methods here! */
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
