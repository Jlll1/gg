package utils

import "strings"

func FormatUsage(command string, args []string, description string) string {
	return "\n  " + strings.Join(append(args, command), ", ") + "\t" + description
}
