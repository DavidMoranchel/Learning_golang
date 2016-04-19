package main 
import (
	"fmt"
	//news
	"time"
	"strings"
)
/*
	It is used for concurrency
	Split separates chains (which are about to separate ,
	as you 'll separate ) method of strings
	Sleep slows the process in milliseconds
	The keyword causes a gorutina go believed
	I mean that several actions are executed at the same time
*/

func main() {
	go name_slow("David")
	fmt.Println("Que aburrido")	
	var wait string
	fmt.Scanln(&wait)
}

func name_slow(name string) {
	letters := strings.Split(name,"")
	for _,letter := range(letters){
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}
}