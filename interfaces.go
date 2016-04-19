package main 
import "fmt"

//interface
type User interface{
	Permissions() int
	Name() string
}

//structs
type Admin struct{
	name string
}
type Editor struct{
	name string
}


//Methods Admin
func (this Admin) Permissions() int{
	return 5
}
func (this Admin) Name() string{
	return this.name
}


//Methods Editor
func (this Editor) Permissions() int{
	return 3
}
func (this Editor) Name() string{
	return this.name
}


func auth(user1 User) string{
	permissions := user1.Permissions()
	if permissions > 4{
		return user1.Name() + " you have permissions of Admin"
	}else if permissions == 3{
		return user1.Name() + " you have permissions of Editor"
	}else{
		return " "
	}
}


func main() {
	admin := Admin{"David"}
	editor := Editor{"Dav"}

	fmt.Println(auth(admin))
	fmt.Println(auth(editor))

	// Put them in an array
	users := []User{Admin{"Juan"},Editor{"Joshua"}}
	for _,user := range users{
		fmt.Println(auth(user))
	}
}