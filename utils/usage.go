package utils

import "strings"

func FormatUsage(command string, args []string, description string) string {
	return "  " + strings.Join(append([]string{command}, args...), ", ") + "\t" + description
}
