package main

import (
	"fmt"
	"os"

	"github.com/Laellekoenig/navigate/internal/navigate/config"
	"github.com/Laellekoenig/navigate/internal/navigate/find"
	"github.com/Laellekoenig/navigate/internal/navigate/fzf"
	"github.com/Laellekoenig/navigate/internal/navigate/tmux"
)

func main() {
	if !config.CheckDependencies() {
		return
	}

	if !tmux.IsTmuxOpen() {
		fmt.Println("Must be in tmux to execute navigate")
		return
	}

	config, err := config.GetConfig()
	if err != nil {
		fmt.Printf("Error when loading config: %s\nDelete the config file if this error repeats.\n", err)
		os.Exit(1)
	}

	res, err := find.FindDirsInDirs(config.NavDirs, 1)
	if err != nil {
		fmt.Printf("Error when finding directories: %s", err)
		return
	}

	if config.IncludeHomeDir && config.HomeDir != "" {
		res = append(res, config.HomeDir)
	}

	selectedDir, err := fzf.GetUserSelection(res)
	if err != nil {
		fmt.Printf("Error when getting user selection: %s", err)
		return
	}

	if selectedDir == "" {
		return
	}

	sessionName := tmux.GetSessionNameFromPath(selectedDir)

	if !tmux.SessionExists(sessionName) {
		err := tmux.CreateSession(sessionName, selectedDir)
		if err != nil {
			fmt.Printf("Error when creating tmux session: %s", err)
			return
		}
	}

	err = tmux.SwitchSession(sessionName)
	if err != nil {
		fmt.Printf("Error when switching tmux session: %s", err)
		return
	}
}
