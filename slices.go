package main 

import "fmt"

/* 
   It is based on arrangements
   They are scalable
   Do not set the size at runtime
   Nullo they have value if you do not declare
   3 values:
   1. Pointer to the array
   2. Length under which aims
   3. Capacity
*/

func main() {

	//Declaration
	var slices []int
	fmt.Println(slices)
	slices1 := []int{1,2,3,4}
	fmt.Println(slices1)

	//length
	fmt.Println(len(slices1))

	//Array to slice
	arreglo := [3]int{1,2,3}
	//slicing(from an array)
	slice := arreglo[:2]
	fmt.Println(slice)
	
}