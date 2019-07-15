package entities

import (
	"database/sql"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
)

func PutEntityInPostgresWithTransaction(entity Entity, pgTxn *sql.Tx) *flaw.Error {
	logger.Debug("entities-put-entity-in-postgres", entity)

	_, err := pgTxn.Exec(
		`INSERT INTO entities (prefix, id, etag, created_at, entity)
     VALUES ($1, $2, $3, $4, $5)`,
		entity.Prefix(),
		entity.ID(),
		entity.Etag(),
		entity.CreatedAt(),
		entity.DBJSON(),
	)

	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"entities_pkey\"" {
			logger.Debug("entities-put-entity-in-postgres-duplicate-key", entity)

			return flaw.New("duplicate entity").Wrap("cannot PutEntityInPostgresWithTransaction")
		}

		return flaw.From(err).Wrap("cannot PutEntityInPostgresWithTransaction")
	}

	return nil
}
