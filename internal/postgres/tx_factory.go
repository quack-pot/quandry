package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quack-pot/quandry/internal/core"
)

type PostgresTx struct {
	tx  pgx.Tx
	ctx context.Context
}

func (tx *PostgresTx) Commit() error {
	return tx.tx.Commit(tx.ctx)
}

func (tx *PostgresTx) Rollback() {
	tx.tx.Rollback(tx.ctx) // Error ignored here since Rollback is error handling from user POV
}

func (db *PostgresDatabase) NewTx(ctx context.Context) (core.IDbTransaction, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	tx, err := db.connections_pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}

	return &PostgresTx{
		tx:  tx,
		ctx: ctx,
	}, nil
}

func (db *PostgresDatabase) WithTx(ctx context.Context, fn func(core.IQueryable) error) error {
	tx, err := db.NewTx(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := fn(tx); err != nil {
		return err
	}

	return tx.Commit()
}
