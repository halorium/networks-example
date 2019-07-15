package main

import (
	"os"

	"github.com/halorium/networks-example/env-vars"
)

func main() {
	_ = os.Setenv("PGUSER", envvars.PostgresUser)
	_ = os.Setenv("PGPASSWORD", envvars.PostgresPassword)

	state := NewState()

	list(state)
}
