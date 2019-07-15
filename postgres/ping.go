package postgres

import (
	"database/sql"
	"time"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
)

func ping(pg *sql.DB) {
	for {
		err := pg.Ping()

		if err != nil {
			logger.Warn("postgres-wait-ping", flaw.From(err).Wrap("cannot ping"))
			time.Sleep(time.Second)
			continue
		}

		break
	}
}
