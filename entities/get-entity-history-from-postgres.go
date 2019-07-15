package entities

import (
	"reflect"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/postgres"
)

func GetEntityHistoryFromPostgres(entity Entity) ([]Entity, *flaw.Error) {
	logger.Debug("entities-get-entity-history-from-postgres", entity)

	pgClient := postgres.ClientFromPool()

	rows, err := pgClient.Query(
		`SELECT entity
     FROM entities
     WHERE prefix = $1
     AND id = $2
     ORDER BY prefix, id, created_at DESC`,
		entity.Prefix(),
		entity.ID(),
	)

	postgres.ClientToPool(pgClient)

	if err != nil {
		return nil, flaw.From(err).Wrap("cannot GetEntityHistoryFromPostgres")
	}

	defer rows.Close()

	collection := make([]Entity, 0)

	for rows.Next() {
		newEntity := reflect.New(reflect.TypeOf(entity).Elem()).Interface().(Entity)

		err = rows.Scan(newEntity)

		if err != nil {
			return nil, flaw.From(err).Wrap("cannot GetEntityHistoryFromPostgres")
		}

		collection = append(collection, newEntity)
	}

	return collection, nil
}
