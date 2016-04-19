package main 
import "fmt"

/*
	Para hacer metodos es necesario declarar una nueva funcion
	en los primeros () va el struct y un calificativo con el
	que nos vamos a referir al metodo, despues nombredefuncion()
	y tipo de dato
*/

type User struct{
	age int
	name, last_name string
}


//hace una copia del objeto user 
func (this User) full_name() string{
	return this.name + " " + this.last_name
}

//Para cambiar atributos se pasa el puntero 
//ES MUCHO MAS EFICIENTE  
func (this *User) set_name(n string){
	this.name = n 
}

func main() {
	david := User{19,"David","Cermeño"}	
	fmt.Println(david.full_name())
	david.set_name("Daniel")
	fmt.Println(david.name)
	fmt.Println(david.full_name())
}