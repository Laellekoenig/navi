package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Laellekoenig/navi/internal/navi/config"
	"github.com/Laellekoenig/navi/internal/navi/find"
	"github.com/Laellekoenig/navi/internal/navi/fzf"
	"github.com/Laellekoenig/navi/internal/navi/ssh"
	"github.com/Laellekoenig/navi/internal/navi/tmux"
)

func main() {
	var list bool
	var selectedDir string
	flag.BoolVar(&list, "list", false, "List all available directories")
	flag.StringVar(&selectedDir, "select", "", "Open session of selected directory")
	flag.Parse()

	if !config.CheckDependencies() {
		return
	}

	if !tmux.IsTmuxOpen() {
		fmt.Println("Must be in tmux to execute navi")
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

	if len(config.SshConnections) > 0 {
		ssh.AddSshOptions(&res, config)
	}

	if list {
		for _, dir := range res {
			fmt.Println(dir)
		}
		return
	}

	if selectedDir == "" {
		selectedDir, err = fzf.GetUserSelection(res)
		if err != nil {
			fmt.Printf("Error when getting user selection: %s", err)
			return
		}
	}

	if selectedDir == "" {
		return
	}

	var sessionName string
	isSsh, connection := ssh.IsSshOption(&selectedDir, config)
	if isSsh {
		sessionName = ssh.GetSessionName(&selectedDir)
	} else {
		sessionName = tmux.GetSessionNameFromPath(selectedDir)
	}

	if !tmux.SessionExists(sessionName) {
		if isSsh {
			err = tmux.CreateSshSession(sessionName, connection)
		} else {
			err = tmux.CreateSession(sessionName, selectedDir)
		}

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
