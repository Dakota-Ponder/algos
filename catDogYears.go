package main

func CalculateYears(years int) (result [3]int) {
	var catYears, dogYears int
  switch  {
  case years == 1:
	catYears = 15
	dogYears = 15
  case years == 2:
	catYears = 15 + 9
	dogYears = 15 + 9
  default:
	catYears = 15 + 9 + (years - 2)* 4 
	dogYears = 15 + 9 + (years - 2)* 5
  }
  result = [3]int {years, catYears, dogYears}
  return
}