package main

// Generate all of the Fibonacci numbers starting with 1 and 2 and ending
// on the highest number before exceeding the parameter's value
// Sum all of the even numbers you generate and return that int
//
//	1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
func SumEvenFibonacci(limit int) int {

	// base case
	if limit < 2 {
		return 1
	}

	sum := 0
	a, b := 1, 2

	for a <= limit {
		if a%2 == 0 {
			sum += a
		}

		// update a and b to next nums in sequence
		a, b = b, a+b
	}
	return sum
}
