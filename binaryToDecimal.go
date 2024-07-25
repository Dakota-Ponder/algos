package main

import "math"

// function that converts binary to decimal
func BinToDec(bin string) int {
    decimal := 0
    length := len(bin)
    
    // Iterate over each character in the binary string
    for i, ch := range bin {
        // Calculate the position value from right to left
        power := length - i - 1
        
        // If the character is '1', add the corresponding power of 2 to the decimal result
        if ch == '1' {
            decimal += int(math.Pow(2, float64(power)))
        }
    }
    
    return decimal
}