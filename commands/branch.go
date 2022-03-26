package commands

import (
	"fmt"
	"strings"

	"github.com/Jlll1/gg/utils"
)

func BranchUsage() string {
	return utils.FormatUsage("b", []string{}, "Equivalent to git branch")
}

func Branch() string {
	result, _ := Git("branch", []string{})
	return result
}

func BranchesClearUsage() string {
	return utils.FormatUsage("bc", []string{}, "Deletes all local branches, with the exceptions of 'main' and 'master'")
}

func BranchesClear() string {
	var branchesToDelete []string
	for _, branch := range strings.Split(Branch(), "\n") {
		if branch == "" || branch[0] == '*' {
			continue
		}

		branch = strings.Trim(branch, " ")
		if branch != "main" && branch != "master" {
			branchesToDelete = append(branchesToDelete, branch)
		}
	}

	Git("branch", append([]string{"-D"}, branchesToDelete...))

	if len(branchesToDelete) > 0 {
		return fmt.Sprintf("%d branches were deleted:\n  ", len(branchesToDelete)) +
			strings.Join(branchesToDelete, "\n  ")
	} else {
		return "No branches to delete!"
	}
}
