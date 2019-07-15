package sleep

import (
	"math/rand"
	"time"
)

func ForAwhile(duration time.Duration) {
	time.Sleep(
		time.Duration(
			rand.Float32() * float32(duration),
		),
	)
}
