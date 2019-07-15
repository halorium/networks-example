package entities

import (
	"database/sql"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/postgres"
)

func GetEntityFromPostgres(entity Entity) *flaw.Error {
	logger.Debug("entities-get-entity-from-postgres", entity)

	pgClient := postgres.ClientFromPool()

	err := pgClient.QueryRow(
		`SELECT entity
     FROM entities
     WHERE prefix = $1
     AND id = $2
     ORDER BY created_at DESC
     LIMIT 1`,
		entity.Prefix(),
		entity.ID(),
	).Scan(entity)

	postgres.ClientToPool(pgClient)

	if err != nil {
		if err == sql.ErrNoRows {
			return flaw.New("not found").Wrap("cannot GetEntityFromPostgres")
		}

		return flaw.From(err).Wrap("cannot GetEntityFromPostgres")
	}

	return nil
}
