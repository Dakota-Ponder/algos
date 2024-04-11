package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 99, 100}
	// val := 60

	// found := BinarySearch(arr, val)
	reverse := ReverseList(arr)

	fmt.Println(reverse)
}
