package main

import "unicode"

func ReverseLetters(s string) string {
	word := []rune{} // map of runes
	for _, c := range s {
		if unicode.IsLetter(c) {
			word = append([]rune{c}, word...)
		}
	}
	return string(word)
}
