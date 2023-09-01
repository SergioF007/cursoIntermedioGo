package main

import "fmt"

type Employee struct {
	id int
}

type Person struct {
	name string
	age  int
}

type FullTimeEmployee struct {
	Employee
	Person
}

type TemporaryEmployee struct {
	Employee
	Person
	taxRate int
}

// Creamos la interface
type PrintInfo interface {
	getMessage() string // instacio el metodo
}

// Como en Go no tenemos una palabra reservada para poder implemetar una interfaz lo tenemos que hacer de la siguiente manera
// creamos un metodo que me va a resivir el objeto y nombramos el metodo igual como esta definido en la interfaz
func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full time employee"
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

// para poder implemetar el metodo de la interfaz creamos
// un funcion con el mismo nombre que me resiva como parametro la interfaz
// ya que este metodo es el que vamos a implemetar en main para que me ejecute la interfaz
func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())

}

func main() {

	ftEmployee := FullTimeEmployee{}
	tEmployee := TemporaryEmployee{}

	getMessage(ftEmployee)
	getMessage(tEmployee)
}
