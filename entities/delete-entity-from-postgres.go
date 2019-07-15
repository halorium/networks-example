package entities

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/postgres"
)

func DeleteEntityFromPostgres(entity Entity) *flaw.Error {
	logger.Debug("entities-delete-entity-from-postgres", entity)

	likePattern := entity.Prefix() + "/" + entity.ID() + "/%"

	var subordinateCount int64

	pgClient := postgres.ClientFromPool()
	defer postgres.ClientToPool(pgClient)

	pgTxn, err := pgClient.Begin()

	if err != nil {
		return flaw.From(err).Wrap("cannot DeleteEntityFromPostgres")
	}

	pgTxn.QueryRow(
		`SELECT COUNT(*)
     FROM entities
     WHERE prefix LIKE $1`,
		likePattern,
	).Scan(&subordinateCount)

	if subordinateCount != 0 {
		err := pgTxn.Rollback()

		if err != nil {
			return flaw.From(err).Wrap("cannot DeleteEntityFromPostgres")
		}

		return flaw.New("subordinate entities exist").Wrap("cannot DeleteEntityFromPostgres")
	}

	_, err = pgTxn.Exec(
		`DELETE FROM entities
     WHERE prefix = $1
     AND id = $2`,
		entity.Prefix(),
		entity.ID(),
	)

	if err != nil {
		rollbackErr := pgTxn.Rollback()

		if rollbackErr != nil {
			return flaw.From(rollbackErr).Wrap("cannot DeleteEntityFromPostgres")
		}

		return flaw.From(err).Wrap("cannot DeleteEntityFromPostgres")
	}

	err = pgTxn.Commit()

	if err != nil {
		return flaw.From(err).Wrap("cannot DeleteEntityFromPostgres")
	}

	return nil
}
