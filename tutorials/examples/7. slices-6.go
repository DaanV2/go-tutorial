package examples

import "slices"

func get(item []int) []int {
	index := slices.Index(item, 1)
	if index == -1 {
		return item
	}

	return item[index:]
}