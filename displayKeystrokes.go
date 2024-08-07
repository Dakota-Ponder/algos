package main

// need to refactor the map to make using strings package and contains
func MobileKeyboard(str string) int {
	keystroke := map[rune]int{
		'1': 1,
		'2': 1,
		'3': 1,
		'4': 1,
		'5': 1,
		'6': 1,
		'7': 1,
		'8': 1,
		'9': 1,
		'0': 1,
		'a': 2,
		'b': 3,
		'c': 4,
		'd': 2,
		'e': 3,
		'f': 4,
		'g': 2,
		'h': 3,
		'i': 4,
		'j': 2,
		'k': 3,
		'l': 4,
		'm': 2,
		'n': 3,
		'o': 4,
		'p': 2,
		'q': 3,
		'r': 4,
		's': 5,
		't': 2,
		'u': 3,
		'v': 4,
		'w': 2,
		'x': 3,
		'y': 4,
		'z': 5,
		'*': 1,
		'#': 1,
	}

	// variable to hold the number of strokes to reach letter
	strokes := 0

	// loop through the string
	for _, char := range str {

		// add the value of the char from the map to the total
		strokes += keystroke[char]
	}

	return strokes
}
