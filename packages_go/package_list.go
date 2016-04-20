package main 
import (
	"fmt"
	"container/list"
)

func main() {
	//create a new list
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(12, e4)
	l.InsertAfter(0, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}