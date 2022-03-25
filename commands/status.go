package commands

import (
	"os/exec"

	"github.com/Jlll1/gg/utils"
)

func StatusUsage() string {
	return utils.FormatUsage("s", []string{}, "Prints short status")
}

func Status() string {
	output, _ := exec.Command("git", "status", "--sort").Output()
	return string(output[:])
}
