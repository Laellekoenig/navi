package config

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"
)

type Config struct {
	NavDirs        []string `json:"directories"`
	IncludeHomeDir bool     `json:"includeHomeDir"`
	HomeDir        string
}

func getStandardConfig() *Config {
	homeDir, err := getHomeDir()
	if err != nil {
		return &Config{
			NavDirs:        []string{},
			IncludeHomeDir: true,
			HomeDir:        homeDir,
		}
	}
	return &Config{
		NavDirs:        []string{path.Join(homeDir, ".config")},
		IncludeHomeDir: true,
		HomeDir:        homeDir,
	}
}

func GetConfig() (*Config, error) {
	homeDir, err := getHomeDir()
	if err != nil {
		return nil, err
	}

	fileName := path.Join(homeDir, ".config", "navigate", "config.json")
	file, err := os.Open(fileName)
	if err != nil {
		stdConfig := getStandardConfig()
		err = stdConfig.save()
		if err != nil {
			return nil, err
		}
		return stdConfig, nil
	}
	defer file.Close()

	config := &Config{}
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, err
	}
	config.HomeDir = homeDir

	return config, nil
}

func (c *Config) save() error {
	homeDir, err := getHomeDir()
	if err != nil {
		return err
	}

	configDir := path.Join(homeDir, ".config", "navigate")
	_, err = os.Stat(configDir)
	if err != nil {
		os.MkdirAll(configDir, 0755)
	}

	jsonData, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}

	file, err := os.Create(path.Join(configDir, "config.json"))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(string(jsonData))
	if err != nil {
		return err
	}

	return nil
}

func getHomeDir() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return user.HomeDir, nil
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
