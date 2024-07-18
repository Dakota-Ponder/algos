package main

// Count the number of occurrences of each character and return it as a (list of tuples)
// in order of appearance.
// For empty output return (an empty list).

type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(input string) []Tuple {
	// Map to keep track of character counts
	countMap := make(map[rune]int)
	// Slice to keep track of the order of appearance
	orderSlice := []rune{}

	for _, char := range input {
		// If the character appears for the first time
		if _, exists := countMap[char]; !exists {
			orderSlice = append(orderSlice, char)
		}
		// Count occurrences
		countMap[char]++
	}

	// Preparing the result as a slice of Tuples
	result := []Tuple{}
	for _, char := range orderSlice {
		result = append(result, Tuple{Char: char, Count: countMap[char]})
	}

	return result
}
