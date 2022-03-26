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
	fmt.Fprintln(writer, commands.StatusUsage())
	fmt.Fprintln(writer, commands.AddCommitUsage())
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
	case "s":
		result = commands.Status()
	case "ac":
		result = commands.AddCommit(os.Args[2:])
	default:
		result = usage()
	}

	fmt.Printf(result + "\n")
}
