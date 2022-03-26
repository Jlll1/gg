package commands

import (
	"github.com/Jlll1/gg/utils"
)

func StatusUsage() string {
	return utils.FormatUsage("s", []string{}, "Equivalent to git status --short")
}

func Status() string {
	output, _ := Git("status", []string{"--short"})
	return output
}
