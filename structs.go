package main 
import "fmt"

/*
A struct is declared outside the main and within
They specify all the variables that take
Parabra reserved the type used to identify
type in this case would type ' list '
The structures are mutable ( can be changed value )
*/

type list struct{
	age int
	name string
	married bool 
}

type list2 struct{
	age int
	sisters int 
	name string
}

func main() {

	var l list
	fmt.Println(l)

	//or you can define so

	l1 := list{age: 19,name: "David", married: true}
	fmt.Println(l1)
	//So accessed parameters
	fmt.Println(l1.name)

	//You can also declare by order

	l2 := list2{19,2,"Juan"}
	fmt.Println(l2)

	//new
	l3 := new(list)
	fmt.Println(l3)



}