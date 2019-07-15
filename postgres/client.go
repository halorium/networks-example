package postgres

import (
	"database/sql"

	// stdlib made me do it!
	_ "github.com/lib/pq"
)

type Client struct {
	*sql.DB
}
