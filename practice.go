package main

func ReverseList(lst []int) []int {
	newSlice := make([]int, len(lst)) // make a new slice with the same length as the old slice

	// loop through the slice starting from the last element in the array
	for i := range lst {
		newSlice[len(lst)-1-i] = lst[i]
	}

	return newSlice
}
