package main 
import "fmt"

func main() {
	arr1 := []int{1,2,3,4}
	c := make([]int,len(arr1),cap(arr1)*2)
	fmt.Println(arr1,c)
/*
	copy (where you are copying ,
	you are going to copy)
	The copy function copies the minimum elements
	If the array is 0, copied 0 items
 */
	copy(c,arr1)
	fmt.Println(arr1,c)
}