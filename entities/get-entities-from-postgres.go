package entities

import (
	"reflect"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/postgres"
)

func GetEntitiesFromPostgres(entity Entity) ([]Entity, *flaw.Error) {
	logger.Debug("entities-get-entities-from-postgres", nil)

	pgClient := postgres.ClientFromPool()

	rows, err := pgClient.Query(
		`SELECT DISTINCT ON (prefix, id) entity
     FROM entities
     WHERE prefix = $1
     ORDER BY prefix, id, created_at DESC`,
		entity.Prefix(),
	)

	postgres.ClientToPool(pgClient)

	if err != nil {
		return nil, flaw.From(err).Wrap("cannot GetCollectionFromPostgres")
	}

	defer rows.Close()

	collection := make([]Entity, 0)

	for rows.Next() {
		newEntity := reflect.New(reflect.TypeOf(entity).Elem()).Interface().(Entity)

		err = rows.Scan(newEntity)

		if err != nil {
			return nil, flaw.From(err).Wrap("cannot GetCollectionFromPostgres")
		}

		collection = append(collection, newEntity)
	}

	return collection, nil
}
