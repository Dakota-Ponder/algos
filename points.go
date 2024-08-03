package main

import (
	"strconv"
	"strings"
)

/* 
Points are awarded for each match as follows:

    if x > y: 3 points (win)
    if x < y: 0 points (loss)
    if x = y: 1 point (tie)

write a function that takes this collection and returns the number of points our team (x) 
got in the championship by the rules given above.
*/
func Points(games []string) int {

	totalPoints := 0

  for _, game := range games{
	scores := strings.Split(game, ":")
	x,_ := strconv.Atoi(scores[0])
	y,_ := strconv.Atoi(scores[1])

	if x > y{
		totalPoints += 3
	} else if x == y {
		totalPoints += 1
	} 
  }


  return totalPoints
}