package main

import (
	"fmt"

	"github.com/Laellekoenig/navigate/pkg/navigate/config"
	"github.com/Laellekoenig/navigate/pkg/navigate/find"
	"github.com/Laellekoenig/navigate/pkg/navigate/fzf"
	"github.com/Laellekoenig/navigate/pkg/navigate/tmux"
)

func main() {
	if !config.CheckDependencies() {
		return
	}

	if !tmux.IsTmuxOpen() {
		fmt.Println("Must be in tmux to execute navigate")
		return
	}

	navDirs, err := config.GetNavDirs()
	if err != nil {
		panic(err)
	}

	res, err := find.FindDirsInDirs(navDirs, 1)
	if err != nil {
		fmt.Printf("Error when finding directories: %s", err)
		return
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
