package flaw

import (
	"runtime"
	"strings"
)

func getFrames() []frame {
	frames := []frame{}

	atTop := true

	for i := 1; ; i++ {
		_, pathname, line, ok := runtime.Caller(i)

		if !ok {
			break
		}

		if atTop && strings.Contains(pathname, "networks-example/flaw/") {
			continue
		}

		atTop = false

		frm := frame{
			Pathname: stripPathname(pathname),
			Line:     line,
		}

		frames = append(frames, frm)
	}

	// remove go runtime entrypoints
	return frames[0 : len(frames)-2]
}
