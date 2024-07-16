package main

import (
	"strconv"
)

func Encode(str string, key int) []int {
	// convert the key into a string ("1939")
	keyStr := strconv.Itoa(key)

	// create a slice to hold the keys individual digits  [0, 0, 0, 0]
	keyDigits := make([]int, len(keyStr))

	// loop through each char in keyStr and convert to an int
	// this makes keyDigits equal [1,9,3,9]
	for i, char := range keyStr {
		keyDigits[i], _ = strconv.Atoi(string(char))
	}

	// create a slice to hold the final encoded numbers
	// len of the slice is the same as the len of the input string
	encoded := make([]int, len(str))

	// loop through the string and its digits
	for i, char := range str {
		letterNum := letterToNumber(char) // convert char to its number
		// get the appropriate key digit, use modulo to loop over keyDigits repeatedly
		keyNum := keyDigits[i%len(keyDigits)]
		encoded[i] = letterNum + keyNum
	}
	return encoded
}

func letterToNumber(letter rune) int {
	return int(letter - 'a' + 1)
}
