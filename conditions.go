package main 
import "fmt"

func main() {
	x := 10
	x1 := 10
	y := 20
	z := 30
	
	if(x == x1){
		fmt.Println("x and x1 are equals")
	}

	/*
		They may or may not have brackets
	*/

	if x > y{
		fmt.Println("Nooot")
	}else if x > z{
		fmt.Println("Nooooot")
	}else{
		fmt.Println("x is the most small")
	}
}