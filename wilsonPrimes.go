package main

import "math/big"

/*
Wilson primes satisfy the following condition. Let P represent a prime number.

Then,

((P-1)! + 1) / (P * P)

should give a whole number.
*/

// Function to calculate factorial of a number using recursion
func factorial(n int) *big.Int {
	if n == 0 || n == 1 {
		return big.NewInt(1)
	}

	result := big.NewInt(int64(n))
	return result.Mul(result, factorial(n-1))
}

// Function to check if a number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func AmIWilson(n int) bool {
	if !isPrime(n) {
		return false
	}

	fact := factorial(n - 1)
	fact.Add(fact, big.NewInt(1))

	nBig := big.NewInt(int64(n))
	nSquare := new(big.Int).Mul(nBig, nBig)

	result := new(big.Int)
	result.DivMod(fact, nSquare, result) // Perform (factorial(n-1) + 1) / (n*n)

	return result.Cmp(big.NewInt(0)) == 0
}
