package logger

type entry struct {
	Tag          string      `json:"tag,omitempty"`
	Severity     string      `json:"severity,omitempty"`
	GitBranch    string      `json:"git-branch,omitempty"`
	GitSha       string      `json:"git-sha,omitempty"`
	ErrorMessage string      `json:"error-message,omitempty"`
	Message      interface{} `json:"message,omitempty"`
	ansiPrefix   string
}
