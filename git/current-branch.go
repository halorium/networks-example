package git

import (
	"os/exec"
	"strings"

	"github.com/halorium/networks-example/flaw"
)

func CurrentBranch() (string, *flaw.Error) {
	branchNameBytes, err := exec.Command("git/current-branch").Output()

	if err != nil {
		return "", flaw.From(err).Wrap("cannot CurrentBranch")
	}

	return strings.TrimSpace(string(branchNameBytes)), nil
}
