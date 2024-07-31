package main


/*
 build a program that takes a value, integer , and returns a list of 
 its multiples up to another value, limit . If limit is a multiple of integer, 
 it should be included as well.

FindMultiples(2, 6) -> [2,4,6]
*/

func FindMultiples(integer, limit int) []int {
	// declare a slice 
   var multiples []int

  // loop from the integer to the limit 
  for i := integer; i <= limit; i += integer{
	  multiples = append(multiples, i)

  }
  return multiples
}
