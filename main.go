package main
import "fmt"
func main(){
	arr := [5]int{10,20,30,40,50}
	val := 50

	found := LinearSearch(arr, val)

	fmt.Println("Found at position", found)
}
