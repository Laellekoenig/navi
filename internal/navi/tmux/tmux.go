package tmux

import (
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/Laellekoenig/navi/internal/navi/config"
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

func CreateSshSession(sessionName string, connection *config.Ssh) error {
	var cmd *exec.Cmd

	if connection.KeyPath == "" {
		cmd = exec.Command("tmux", "new-session", "-d", "-s", sessionName, "ssh "+connection.Target)
	} else {
		cmd = exec.Command("tmux", "new-session", "-d", "-s", sessionName, "ssh -i"+connection.KeyPath+" "+connection.Target)
	}

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
