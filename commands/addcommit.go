package commands

import "github.com/Jlll1/gg/utils"

func AddCommitUsage() string {
	return utils.FormatUsage(
		"ac",
		[]string{"[<files>]", "<message>"},
		"Adds files which names contain specified strings to a commit with specified message.\n\t"+
			"Use /<file> to include directory names in search.")
}

func AddCommit(args []string) string {
	if len(args) < 2 {
		return AddCommitUsage()
	}
	return ""

}
