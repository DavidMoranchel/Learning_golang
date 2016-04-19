package main 
import (
	"fmt"
	/*
	import strconv for
	view in https://golang.org/pkg/strconv/
	*/
	"strconv"
)

func main() {
	number := 10
	str := "100"

	
	str_to_number,_ := strconv.Atoi(str)
	fmt.Println(str_to_number + number)

	number_to_string := strconv.Itoa(number)
	fmt.Println("now I am string", number_to_string)
}