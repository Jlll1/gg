package commands

import (
	"github.com/Jlll1/gg/utils"
)

func StatusUsage() string {
	return utils.FormatUsage("s", []string{}, "Prints short status")
}

func Status() string {
	output, _ := Git("status", []string{"--short"})
	return output
}
