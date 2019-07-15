package entities

import (
	"database/sql"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
)

func GetEntityFromPostgresByBodyEmailWithTransaction(entity Entity, email string, pgTxn *sql.Tx) *flaw.Error {
	logger.Debug("entities-get-entity-from-postgres-by-body-email-with-transaction", nil)

	err := pgTxn.QueryRow(
		`SELECT entity
     FROM entities
     WHERE prefix = $1
     AND entity -> 'body' ->> 'email' = $2
     ORDER BY created_at DESC
     LIMIT 1`,
		entity.Prefix(),
		email,
	).Scan(entity)

	if err != nil {
		if err == sql.ErrNoRows {
			return flaw.New("not found").Wrap("cannot GetEntityFromPostgresByBodyEmailWithTransaction")
		}

		return flaw.From(err).Wrap("cannot GetEntityFromPostgresByBodyEmailWithTransaction")
	}

	return nil
}
