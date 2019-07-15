package entities

import (
	"reflect"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/postgres"
)

func GetEntitiesFromPostgresByCreatedAt(entity Entity, date string) ([]Entity, *flaw.Error) {
	logger.Debug("entities-get-entities-from-postgres-by-created-at", nil)

	pgClient := postgres.ClientFromPool()

	rows, err := pgClient.Query(
		`SELECT entity
     FROM entities
     WHERE prefix = $1
     AND to_char(created_at, 'YYYY-MM-DD') = $2
     ORDER BY prefix, id, created_at DESC`,
		entity.Prefix(),
		date,
	)

	postgres.ClientToPool(pgClient)

	if err != nil {
		return nil, flaw.From(err).Wrap("cannot GetEntitiesFromPostgresByCreatedAt")
	}

	defer rows.Close()

	collection := make([]Entity, 0)

	for rows.Next() {
		newEntity := reflect.New(reflect.TypeOf(entity).Elem()).Interface().(Entity)

		err = rows.Scan(newEntity)

		if err != nil {
			return nil, flaw.From(err).Wrap("cannot GetEntitiesFromPostgresByCreatedAt")
		}

		collection = append(collection, newEntity)
	}

	return collection, nil
}
