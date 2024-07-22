package main

import (
	"sort"
)

/*
   Check if a string contains
    consecutive letters as they appear in the English alphabet
	and if each letter occurs only once.
*/

func Solve(s string) bool {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return s[i] < s[j]
	})

	// loop through the sorted runes
	for i := 0; i < len(runes)-1; i++ {
		// check if next rune is followed by consecutive rune
		if runes[i+1]-runes[i] != 1 {
			return false
		}
	}
	// all checks passed
	return true
}
