package main

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func ReduceFraction(fraction [2]int) [2]int {

	numerator := fraction[0]
	denominator := fraction[1]

	// check if the denom is 0
	if denominator == 0 {
		return [2]int{}
	}

	// find the gcd of the numerator and denominator
	gcdVal := gcd(numerator, denominator)

	// reduce both the numerator and denominator
	reducedNumerator := numerator / gcdVal
	reducedDenominator := denominator / gcdVal

	// return the reduced fraction
	return [2]int{reducedNumerator, reducedDenominator}

}
