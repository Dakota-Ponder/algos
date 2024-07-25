package main

import (
	"fmt"
)

func main() {
	binaryStr := "10"
    decimal := BinToDec(binaryStr)
    fmt.Printf("The decimal representation of %s is %d\n", binaryStr, decimal)
}
