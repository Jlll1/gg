package commands

import (
	"fmt"
	"strings"

	"github.com/Jlll1/gg/utils"
)

func FileResetUsage() string {
	return utils.FormatUsage(
		"fr",
		[]string{"<file>"},
		"Equivalent to git checkout -- <file>.")
}

func FileReset(args []string) string {
	if len(args) < 1 {
		return FileResetUsage()
	}

	filePredicate := args[0]
	allFiles := strings.Split(Status(), "\n")

	matchedFiles := utils.Filter(allFiles, func(x string) bool {
		return strings.Contains(x, filePredicate)
	})

	matchedFiles = utils.Map(matchedFiles, func(x string) string {
		return strings.Split(x, " ")[2]
	})

	Git("checkout", []string{"--", fmt.Sprintf("*%s*", filePredicate)})

	return "Files reset:\n" + strings.Join(matchedFiles, "\n")
}
