package main 
import "fmt"

/*
	1 is a memory address
	2. We have no value , we address
	the value
	3. The value zero is nill
*/

func main() {
	var x,y *int
	num := 5
	x = &num
	y = &num

	//Print your memory space
	fmt.Println(x)
	fmt.Println(y)


	/* Print your value in memory
 	Accessed through the operator * */
	fmt.Println(*x)
	fmt.Println(*y)



	*x = 6

	fmt.Println(x)
	fmt.Println(y)

	/* If we change the value does not change
	the address in memory */

	fmt.Println(*x)
	fmt.Println(*y)

}