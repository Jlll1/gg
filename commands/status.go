package commands

import "os/exec"

func StatusUsage() string {
	return "s\tPrints short status"
}

func Status() string {
	output, _ := exec.Command("git", "status", "--sort").Output()
	return string(output[:])
}
