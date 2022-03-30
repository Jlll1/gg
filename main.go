package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/Jlll1/gg/commands"
	"github.com/Jlll1/gg/utils"
)

func usage() string {
	var builder strings.Builder
	builder.WriteString("usage: gg <command> [<args>]\n\n")
	builder.WriteString("List of commands:\n")

	writer := tabwriter.NewWriter(&builder, 0, 8, 2, '\t', tabwriter.AlignRight)
	fmt.Fprintln(writer, commands.AddCommitUsage())
	fmt.Fprintln(writer, commands.BranchUsage())
	fmt.Fprintln(writer, commands.BranchesClearUsage())
	fmt.Fprintln(writer, commands.FileResetUsage())
	fmt.Fprintln(writer, commands.StatusUsage())
	writer.Flush()

	return builder.String()
}

func main() {
	var result string

	command, err := utils.TryGet(os.Args, 1)
	if err != nil {
		fmt.Println(usage())
		return
	}

	switch command {
	case "ac":
		result = commands.AddCommit(os.Args[2:])
	case "b":
		result = commands.Branch()
	case "bc":
		result = commands.BranchesClear()
	case "fr":
		result = commands.FileReset(os.Args[2:])
	case "s":
		result = commands.Status()
	default:
		result = usage()
	}

	fmt.Printf(result + "\n")
}
