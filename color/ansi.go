package color

import (
	"bytes"

	"github.com/mgutz/ansi"
)

var (
	// Reset ANSI sequence
	Reset = ansi.ColorCode("reset")

	// Black ANSI sequence
	Black = ansi.ColorCode("black")
	// LightBlack ANSI sequence
	LightBlack = ansi.ColorCode("black+h")
	// Red ANSI sequence
	Red = ansi.ColorCode("red")
	// LightRed ANSI sequence
	LightRed = ansi.ColorCode("red+h")
	// Green ANSI sequence
	Green = ansi.ColorCode("green")
	// LightGreen ANSI sequence
	LightGreen = ansi.ColorCode("green+h")
	// Yellow ANSI sequence
	Yellow = ansi.ColorCode("yellow")
	// LightYellow ANSI sequence
	LightYellow = ansi.ColorCode("yellow+h")
	// Blue ANSI sequence
	Blue = ansi.ColorCode("blue")
	// LightBlue ANSI sequence
	LightBlue = ansi.ColorCode("blue+h")
	// Magenta ANSI sequence
	Magenta = ansi.ColorCode("magenta")
	// LightMagenta ANSI sequence
	LightMagenta = ansi.ColorCode("magenta+h")
	// Cyan ANSI sequence
	Cyan = ansi.ColorCode("cyan")
	// LightCyan ANSI sequence
	LightCyan = ansi.ColorCode("cyan+h")

	// BlackOnWhite ANSI sequence
	BlackOnWhite = ansi.ColorCode("black:white+h")
)

func Colorize(ansiPrefix string, text []byte) *bytes.Buffer {
	buffer := bytes.NewBufferString(ansiPrefix)

	_, err := buffer.Write(text)

	if err != nil {
		panic(err)
	}

	_, err = buffer.WriteString(Reset)

	if err != nil {
		panic(err)
	}

	return buffer
}

func Discolorize(ansiPrefix string, text []byte) *bytes.Buffer {
	buffer := bytes.NewBuffer(nil)

	_, err := buffer.Write(text)

	if err != nil {
		panic(err)
	}

	return buffer
}
