package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

// crear el metodo Set de Id
func (e *Employee) SetId(id int) {
	e.id = id
}

// crear el metodo Set de Name
func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

// implementado la funcion de la forma 4 de represetar contructores para la creacion de objetos en Go
func NewEmployee(id int, name string, vacation bool) *Employee {

	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}

}

func main() {

	e := Employee{}
	//fmt.Printf("%v", e)

	e.id = 1
	e.name = "Sergio"
	//fmt.Printf("%v", e)

	e.SetId(2) // modifico el Id
	e.SetName("Paula")
	//fmt.Printf("%v", e)
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())

	//Forma 2 de crear objetos
	e2 := Employee{
		id:       1,
		name:     "Pacho",
		vacation: true,
	}

	fmt.Printf("%v", e2)

	// Froma 3 de crear objetos debo de usar el *e3 para acceder a la inf
	// sino solo me muestra una referencia de e3
	e3 := new(Employee)
	fmt.Printf("%v", *e3)

	// Forma 4 - Es la recomendada ya que me da mucha mas flexibilidad
	e4 := NewEmployee(3, "PruebaFuncConst", true)
	fmt.Printf("%v", e4)

}
