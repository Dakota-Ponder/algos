package main

// function that removes the first and last char from a string 
func RemoveChar(word string) string {
  runes := []rune(word)  // turn the string into runes 

  // if length of the string is less than 2, return an empty string 
  if len(runes) <= 2{
	return ""
  }
 // empty string to hold new word 
 st := "" 

  // remove the first and last rune from the string
  for i := 1; i < len(runes) - 1; i++{
	st += string(runes[i])
  } 
  return st
}