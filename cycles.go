package main 
import "fmt"

/*
	In golang for is the only cycle
*/

func main() {
	for i:=0;i<=10;i++{
		fmt.Println(i)
	}

	x:=0
	for x <= 15 {
		fmt.Println(x)
		x++
		if x == 10{
			break
		}
	}





	y:= 0
	for {

		if y == 2{
			y++
			continue
		}
		fmt.Println(y)
		y++

		if y == 10{
			break
		}
	}
}