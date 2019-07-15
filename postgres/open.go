package postgres

import (
	"database/sql"
	"time"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
)

func open() *sql.DB {
	for {
		pg, err := sql.Open("postgres", "")

		if err != nil {
			logger.Warn("postgres-wait-open", flaw.From(err).Wrap("cannot open"))
			time.Sleep(time.Second)
			continue
		}

		return pg
	}
}
