package main 

import "fmt"

/*

Data type has two parameters
[Number of elements ( not modified ) ] ,
data type ( can not be conbinar data types)

*/


func main() {

	/* No initial values */
	var arr [10]string
	fmt.Println(arr)

	/* with initial values */
	arr1 := [3]int{1,2,3}
	fmt.Println(arr1)
	

	/* It takes values ​​if they have not declared */
	arr2 := [4]int{1,2,3}
	fmt.Println(arr2)

	
	/* Size */
	fmt.Println(len(arr2))

	/* Change the value */
	arr2[3] = 15
	fmt.Println(arr2)


	/* iterating */	
	for i:=0;i<len(arr2);i++{
		fmt.Println(arr2[i])
	}

	/* two dimensional arrays */

	var arr3[3][4] int 
	fmt.Println(arr3)
	arr3[0][3] = 4
	fmt.Println(arr3)
}

