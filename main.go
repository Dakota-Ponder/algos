package main

import (
	"fmt"
)

func main() {
	word := "cabbage"  // should return 14 
	score := ScrabbleScore(word)
    fmt.Println(score)
}
