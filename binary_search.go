package main


func BinarySearch(arr[] int, target int) int {
	lo := 0
	hi := len(arr)

	for lo < hi{

		 m := lo + (hi - lo) / 2
		 v := arr[m]

		if v == target{
			return m
		}else if v > target{
			// go to the left side of the array 
			hi = m
		} else {
			// go to the right side of the array 
			lo = m + 1
		}
	}
	return -1

}