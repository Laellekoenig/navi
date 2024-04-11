package fzf

import (
	"os"
	"os/exec"
	"strings"
)

func GetUserSelection(items []string) (string, error) {
	cmd := exec.Command("fzf")
	cmd.Stdin = strings.NewReader(strings.Join(items, "\n"))
	cmd.Stderr = os.Stderr
	selection, err := cmd.Output()

	if exitError, ok := err.(*exec.ExitError); ok {
		if exitError.ExitCode() == 130 {
			return "", nil
		}
	}

	if err != nil {
		return "", err
	}

	// trim newline
	selection = selection[:len(selection)-1]

	return string(selection), nil
}
