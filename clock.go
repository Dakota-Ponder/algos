package main 

// Your task is to write a function which returns the time since midnight in milliseconds.

func Past(h, m, s int) int {
	hour := h * 60 * 60 * 1000
	min :=  m * 60 * 1000
	sec := s * 1000


    return hour + min + sec 
}