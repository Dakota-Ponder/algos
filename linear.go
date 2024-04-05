package main

func LinearSearch(arr [5]int, val int) int {
	for i := 0; i < len(arr); i++{
		if arr[i] == val{
			return i
		}
	}
	return -1
}

