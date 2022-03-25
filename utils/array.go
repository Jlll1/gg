package utils

import "fmt"

func TryGet(arr []string, index int) (string, error) {
	if index < 0 {
		return "", fmt.Errorf("Index %d cannot be smaller than 0", index)
	}

	if index >= len(arr) {
		return "", fmt.Errorf("Tried to access out of range index %d", index)
	}

	return arr[index], nil
}
