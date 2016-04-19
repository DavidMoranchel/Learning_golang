package main
import "fmt"

type Human struct{
	name string
}

type Human1 struct{
	name string
}

// You spend only struct fields you want to share
type Tutor struct{
	Human
}
type Tutor1 struct{
	Human
	Human1
}

// Methods it by going to Human
func (this Human) talk() string{
	return "Hi"
}


// Accessing methods tutor
// Here no ambiguedas
func (this Tutor1) talk1() string{
	return "Hello"
}


func main() {
	tutor := Tutor{Human{"David"}}
	
	// You access the method without having to do tutor.Human.name
	fmt.Println(tutor.name)

	// As there are two anonymous fields you have to
	// Specify where they come
	tutor1 := Tutor1{Human{"Juan"}, Human1{"Daniel"}}
	fmt.Println(tutor1.Human.name)
	fmt.Println(tutor1.Human1.name)
	fmt.Println(tutor1.talk())
	fmt.Println(tutor1.talk1())
}