package entities

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/postgres"
)

func DeleteEntitiesFromPostgres(entity Entity) *flaw.Error {
	logger.Debug("entities-delete-entities-from-postgres", entity)

	likePattern := entity.Prefix() + "/[^/]+/%"

	var subordinateCount int64

	pgClient := postgres.ClientFromPool()
	defer postgres.ClientToPool(pgClient)

	pgTxn, err := pgClient.Begin()

	if err != nil {
		return flaw.From(err).Wrap("cannot DeleteEntitiesFromPostgres")
	}

	pgTxn.QueryRow(
		`SELECT COUNT(*)
     FROM entities
     WHERE prefix SIMILAR TO $1`,
		likePattern,
	).Scan(&subordinateCount)

	if subordinateCount != 0 {
		err := pgTxn.Rollback()

		if err != nil {
			return flaw.From(err).Wrap("cannot DeleteEntitiesFromPostgres")
		}

		return flaw.New("subordinate entities exist").Wrap("cannot DeleteEntitiesFromPostgres")
	}

	_, err = pgTxn.Exec(
		`DELETE FROM entities
     WHERE prefix = $1`,
		entity.Prefix(),
	)

	if err != nil {
		rollbackErr := pgTxn.Rollback()

		if rollbackErr != nil {
			return flaw.From(rollbackErr).Wrap("cannot DeleteEntitiesFromPostgres")
		}

		return flaw.From(err).Wrap("cannot DeleteEntitiesFromPostgres")
	}

	err = pgTxn.Commit()

	if err != nil {
		return flaw.From(err).Wrap("cannot DeleteEntitiesFromPostgres")
	}

	return nil
}
