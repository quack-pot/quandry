package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/quack-pot/quandry/internal/core"
)

type PostgresDatabase struct {
	connections_pool *pgxpool.Pool
}

func NewPostgresDatabase(
	connection_string string,
	options *core.DbOptions,
) (core.IDatabase, error) {
	pgx_config, err := pgxpool.ParseConfig(connection_string)

	if err != nil {
		return nil, err
	}

	pgx_config.MinConns = int32(options.MinConnections)
	pgx_config.MaxConns = int32(options.MaxConnections)
	pgx_config.MaxConnLifetime = time.Duration(options.MaxConnectionLifetime)
	pgx_config.MaxConnIdleTime = time.Duration(options.MaxConnectionIdleTime)
	pgx_config.HealthCheckPeriod = time.Duration(options.HealthCheckPeriod)
	pgx_config.ConnConfig.RuntimeParams["timezone"] = "UTC"

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(options.EstablishConnectionTimeout),
	)
	defer cancel()

	db_pool, err := pgxpool.NewWithConfig(ctx, pgx_config)

	if err != nil {
		return nil, err
	}

	if err := db_pool.Ping(ctx); err != nil {
		db_pool.Close()
		return nil, err
	}

	return &PostgresDatabase{
		connections_pool: db_pool,
	}, nil
}

func (db *PostgresDatabase) Close() error {
	db.connections_pool.Close()
	return nil
}
