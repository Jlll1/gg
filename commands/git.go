package commands

import "os/exec"

func Git(command string, args []string) (string, error) {
	output, err := exec.Command("git", append([]string{command}, args...)...).Output()
	return string(output[:]), err
}
