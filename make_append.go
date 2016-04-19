package main 

import "fmt"

func main() {

/*
	parameters:
	1. Data Type
	2. much data at that time
	3. the ability can have
	the slice

	With append we add a data, the length
	serious 4 with a capacity of 5
	Append extiene capacity when
	exceeds
	Exceeding reduces efficiency
*/

	slice := make([]int,3,5)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice,2)

	fmt.Println(slice)
	fmt.Println(cap(slice))

}