package tmux

import (
	"os"
	"os/exec"
	"path"
	"strings"
)

func IsTmuxOpen() bool {
	return os.Getenv("TMUX") != ""
}

func SessionExists(sessionName string) bool {
	cmd := exec.Command("tmux", "has-session", "-t", sessionName)
	err := cmd.Run()
	return err == nil
}

func CreateSession(sessionName string, dir string) error {
	cmd := exec.Command("tmux", "new-session", "-d", "-s", sessionName, "-c", dir)
	return cmd.Run()
}

func SwitchSession(sessionName string) error {
	cmd := exec.Command("tmux", "switch-client", "-t", sessionName)
	return cmd.Run()
}

func GetSessionNameFromPath(p string) string {
	sessionName := path.Base(p)
	return strings.ReplaceAll(sessionName, ".", "_")
}
