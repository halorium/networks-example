package entities

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/postgres"
)

func PutEntityInPostgres(entity Entity) *flaw.Error {
	logger.Debug("entities-put-entity-in-postgres", entity)

	pgClient := postgres.ClientFromPool()

	_, err := pgClient.Exec(
		`INSERT INTO entities (prefix, id, etag, created_at, entity)
     VALUES ($1, $2, $3, $4, $5)`,
		entity.Prefix(),
		entity.ID(),
		entity.Etag(),
		entity.CreatedAt(),
		entity.DBJSON(),
	)

	postgres.ClientToPool(pgClient)

	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"entities_pkey\"" {
			logger.Debug("entities-put-entity-in-postgres-duplicate-key", entity)

			return nil
		}

		return flaw.From(err).Wrap("cannot PutEntityInPostgres")
	}

	return nil
}
