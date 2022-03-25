package main

import (
	"fmt"
	"os"

	"github.com/Jlll1/gg/commands"
	"github.com/Jlll1/gg/utils"
)

func usage() string {
	return "usage: gg <command> [<args>]" + "\n\n" +
		"List of commands:" + "\n  " +
		commands.StatusUsage()
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
	default:
		result = usage()
	}

	fmt.Printf(result)
}
