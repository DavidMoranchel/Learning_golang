package main 
import "fmt"

func main() {
	var x,y int
	fmt.Println("Insert the value of x")
	/*
		Scanf are the parameters 
		( data type to receive,
		 direction of the variable (pointer)
		)
	*/
	fmt.Scanf("%d",&x)
	fmt.Println("Insert the value of y")
	fmt.Scanf("%d",&y)
	fmt.Println("x =",x,"\ny =",y)
}