package main

func LinearSearch(slice []int, target int) int {
	for index, value := range slice{
		if value == target {
			return index
		}
	}
	return -1
}

