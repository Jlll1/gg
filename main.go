package main

import (
	"fmt"
	"os"

	"github.com/Jlll1/gg/commands"
)

func main() {
	var result string
	command := os.Args[1]

	switch command {
	case "s":
		result = commands.Status()
	}

	fmt.Printf(result)
}
