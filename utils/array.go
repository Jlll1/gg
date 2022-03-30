package utils

import "fmt"

func TryGet[T any](arr []T, index int) (T, error) {
	var result T
	if index < 0 {
		return result, fmt.Errorf("Index %d cannot be smaller than 0", index)
	}

	if index >= len(arr) {
		return result, fmt.Errorf("Tried to access out of range index %d", index)
	}

	return arr[index], nil
}

func Filter[T any](arr []T, predicate func(item T) bool) []T {
	count := 0
	for _, item := range arr {
		if predicate(item) {
			arr[count] = item
			count++
		}
	}

	return arr[:count]
}

func Map[T any, U any](arr []T, mapping func(item T) U) []U {
	result := []U{}
	for _, item := range arr {
		result = append(result, mapping(item))
	}

	return result
}
