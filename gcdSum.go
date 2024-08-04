package main 

/*
Given the sum and gcd of two numbers, return those two numbers in ascending order. 
*/

func SolveGCD(s int, g int) []int {
	// check if s is divisible by g 
	if s % g != 0{
		return []int{-1,-1}
	}

	// Find the values of x and y such that x + y = s / g
	n := s / g
	x := 1 
	y := n - 1 

	// Return the numbers in ascending order
	return []int{g * x, g * y}

}