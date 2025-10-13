package db

import (
	"context"
	"fmt"

	"github.com/go_geofetch/cmd/models"
	"github.com/go_geofetch/generated"
	"github.com/jackc/pgx/v5"
)

func InitDB(ctx context.Context, env *models.EnvModel) (*pgx.Conn, *generated.Queries, error) {
	conn, err := pgx.Connect(ctx, env.DatabaseUrl)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	queries := generated.New(conn)
	return conn, queries, nil

}
