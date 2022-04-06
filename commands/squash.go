package commands

import "github.com/Jlll1/gg/utils"

func SquashUsage() string {
	return utils.FormatUsage(
		"sq",
		[]string{"<target>", "<message>"},
		"Combines commits up to <target> into single commit with message <message>")
}

func Squash(args []string) string {
	if len(args) < 2 {
		return SquashUsage()
	}

	squashUntil := args[0]
	commitMessage := args[1]
	currentBranch := Branch([]string{"--show-current"})
	commonCommit, _ := Git("merge-base", []string{squashUntil, currentBranch})
	Git("reset", []string{commonCommit})
	result, _ := Git("commit", []string{"-a", "-m", commitMessage})

	return result
}
