package database

import (
	"context"
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/p12u/golib/perrors"
)

// CreateEntDriver returns a Postgres driver for ent connected to the provider
// url argument. The returned driver can be used to initialze ent clients in your
// application
func CreateEntDriver(ctx context.Context, url string) (*entsql.Driver, error) {
	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, perrors.Wrap(ctx, err, "failed to connect to database", nil)
	}

	// Create an ent.Driver from `db`.
	return entsql.OpenDB(dialect.Postgres, db), nil
}
