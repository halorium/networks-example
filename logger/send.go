package logger

func send(severity string, ansiPrefix string, tag string, message interface{}) {
	anEntry := &entry{
		ansiPrefix: ansiPrefix,
		GitBranch:  InGitBranch,
		GitSha:     InGitSha,
		Message:    message,
		Severity:   severity,
		Tag:        tag,
	}

	serializeAndWrite(anEntry)
}
