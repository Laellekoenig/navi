package find

import (
	"fmt"
	"os/exec"
)

func FindDirsInDirs(dirs []string, maxDepth int) ([]string, error) {
	res := make([]string, 0)

	for _, dir := range dirs {
		findCmd := exec.Command("find", dir, "-maxdepth", fmt.Sprint(maxDepth), "-type", "d")
		output, err := findCmd.CombinedOutput()
		if err != nil {
			return nil, fmt.Errorf("%s", string(output))
		}

		output = output[:len(output)-1]
		res = append(res, string(output))
	}

	return res, nil
}
