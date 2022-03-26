package commands

import (
	"errors"
	"fmt"

	"github.com/Jlll1/gg/utils"
)

func AddCommitUsage() string {
	return utils.FormatUsage(
		"ac",
		[]string{"[<files>]", "<message>"},
		"Equivalent to git add [*<file>*] && git commit -m <message>")
}

func AddCommit(args []string) string {
	if len(args) < 2 {
		return AddCommitUsage()
	}

	commitMessage := fmt.Sprintf("-m\"%s\"", args[len(args)-1])

	add(args[:len(args)-1])
	status := Status()
	commit, _ := Git("commit", []string{commitMessage})

	return status + "\n" + commit
}

func add(predicates []string) error {
	if len(predicates) < 1 {
		return errors.New("predicates array cannot be empty")
	}

	var addArgs []string

	if predicates[0] == "." {
		addArgs = []string{"."}
	} else {
		for i := range predicates {
			addArgs = append(addArgs, "*"+predicates[i]+"*")
		}
	}

	Git("add", addArgs)

	return nil
}
