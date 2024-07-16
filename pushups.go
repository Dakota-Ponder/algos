package main

/*
He can do ten kata in an hour, but when he makes a mistake, he must do pushups.
These pushups really tire poor Alex out, so every time he does them they take twice as long.
His first set of redemption pushups takes 5 minutes. Create a function, alexMistakes, that takes
two arguments: the number of kata he needs to complete, and the time in minutes he has to complete them.
Your function should return how many mistakes Alex can afford to make.
*/

func AlexMistakes(numberOfKatas, timeLimit int) int {

	pushupTime := 5
	mistakes := 0

	// total time alex needs to complete given # of katas
	timeForKatas := numberOfKatas * 6

	// calc remaining time available for mistakes
	remainingTime := timeLimit - timeForKatas

	// compute how many sets of pushups he can afford within remianing time, with
	// each set takes double the time of the previous one
	for remainingTime >= pushupTime {
		remainingTime -= pushupTime
		mistakes += 1
		pushupTime *= 2 // each set is twice as long
	}

	return mistakes

}
