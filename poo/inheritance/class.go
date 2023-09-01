package main

import "fmt"

type Employee struct {
	id int
}

type Person struct {
	name string
	age  int
}

// En el siguiente Struc vamos a implemetar propiedades tanto de Empoyee y Person
type FullTimeEmployee struct {
	Employee // los estoy poniendo de manera anomima ya que no lo estoy asignando a ningun atributo como serai el caso:
	Person   // employee Employee / person Person

}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d", p.name, p.age)
}

func main() {

	ftEmployee := FullTimeEmployee{}
	ftEmployee.id = 1
	ftEmployee.name = "Prueba implemetando atributos de otros structs"
	ftEmployee.age = 27

	fmt.Printf("%v", ftEmployee)

	// me genera erro por que go no permite implemetar herencia, por lo tanto.
	// puedo aplicarla hasta cierto punto como es el caso de los atributos en en los metodos el comportamiento no se hereda
	//GetMessage(ftEmployee) // Por lo que este tipo de acciones las podemos representar por medio de interfaces

}
