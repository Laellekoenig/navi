package ssh

import (
	"strings"

	"github.com/Laellekoenig/navi/internal/navi/config"
)

func AddSshOptions(options *[]string, config *config.Config) {
	for _, ssh := range config.SshConnections {
		*options = append(*options, "SSH | "+ssh.Target)
	}
}

func IsSshOption(option *string, config *config.Config) (bool, *config.Ssh) {
	if !strings.HasPrefix(*option, "SSH | ") {
		return false, nil
	}

	for _, ssh := range config.SshConnections {
		if strings.HasSuffix(*option, ssh.Target) {
			return true, &ssh
		}
	}

	return false, nil
}

func GetSessionName(option *string) string {
	return "ssh " + strings.ReplaceAll(strings.TrimPrefix(*option, "SSH | "), ".", "_")
}
