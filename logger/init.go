package logger

import (
	"bytes"
	"os"

	"github.com/halorium/networks-example/color"
	"github.com/halorium/networks-example/serializers"
)

var colorizer func(string, []byte) *bytes.Buffer
var serializer func(interface{}) []byte
var InLogDebugMessages string
var InLogColorization string
var InLogSerialization string
var InGitBranch string
var InGitSha string

func init() {
	InLogDebugMessages = os.Getenv("IN_LOG_DEBUG_MESSAGES")
	InLogColorization = os.Getenv("IN_LOG_COLORIZATION")
	InLogSerialization = os.Getenv("IN_LOG_SERIALIZATION")
	InGitBranch = os.Getenv("IN_GIT_BRANCH")
	InGitSha = os.Getenv("IN_GIT_SHA")

	if InLogColorization == "true" {
		colorizer = color.Colorize
	} else {
		colorizer = color.Discolorize
	}

	switch InLogSerialization {
	case "json-pretty":
		serializer = serializers.JSONPrettySerializer
	case "json-compact":
		serializer = serializers.JSONCompactSerializer
	default:
		serializer = serializers.JSONCompactSerializer
	}
}
