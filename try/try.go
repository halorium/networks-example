package try

import (
	"time"

	"github.com/halorium/networks-example/flaw"
)

type Try struct {
	Start   time.Time     `json:"start"`
	Max     uint          `json:"max"`
	Tries   uint          `json:"tries"`
	Elapsed time.Duration `json:"elapsed"`
	Error   *flaw.Error   `json:"error"`
}
