package halt

import "github.com/halorium/networks-example/flaw"

func PanicWith(message string) {
	panic(flaw.New(message))
}
