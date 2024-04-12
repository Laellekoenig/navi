package config

import (
	"fmt"
	"os/exec"
	"os/user"
	"path"
)

func getHomeDir() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return user.HomeDir, nil
}

func GetNavDirs() ([]string, error) {
	homeDir, err := getHomeDir()
	if err != nil {
		return nil, err
	}

	dirs := []string{".config", "code", "uni"}
	for i, dir := range dirs {
		dirs[i] = path.Join(homeDir, dir)
	}

	return dirs, nil
}

func CheckDependencies() bool {
	deps := []string{"fzf", "tmux", "find"}

	for _, dep := range deps {
		_, err := exec.LookPath(dep)
		if err != nil {
			fmt.Printf("Error: %s not found in PATH\n", dep)
			return false
		}
	}

	return true
}
