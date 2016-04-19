package main 
import "fmt"

func main() {
	var x,y int
	fmt.Println("Ingresa el valor de x")
	/*
		Scanf are the parameters 
		( data type to receive,
		 direction of the variable (pointer)
		)
	*/
	fmt.Scanf("%d",&x)
	fmt.Println("Ingresa el valor de y")
	fmt.Scanf("%d",&y)
	fmt.Println("x =",x,"\ny =",y)
}