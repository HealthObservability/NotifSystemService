package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"net"

	"github.com/HealthObservability/NotifSystemService/internal/config"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	//_ "github.com/lib/pq" // postgres
)

type Postgres struct {
	pool *pgxpool.Pool
}

func MustNew(ctx context.Context, cfg config.Postgres) *Postgres {
	log := logger.GetLogger().WithField("op", "postgres.MustNew")
	pool, err := initConn(ctx, cfg)
	if err != nil {
		log.WithError(err).Fatal("error connecting to database")
	}
	log.Info("Connected to postgres successfully")

	return &Postgres{pool}
}

func initConn(ctx context.Context, cfg config.Postgres) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.User, cfg.Password, net.JoinHostPort(cfg.Host, cfg.Port), cfg.Database,
	)
	
	dbPool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, err
	}

	if err := dbPool.Ping(ctx); err != nil {
		return nil, err
	}

	return dbPool, nil
}

func (p *Postgres) Shutdown(ctx context.Context) error {
	p.pool.Close()
	return nil // fixme
}
