package commands

import "os/exec"

func Status() string {
	output, _ := exec.Command("git", "status", "-s").Output()
	return string(output[:])
}
