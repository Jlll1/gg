package gg

import (
	"os"
	"os/exec"
)

func main() {
	command := os.Args[0]
	exec.Command("git", command)
}
