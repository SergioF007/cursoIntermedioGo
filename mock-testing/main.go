package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

// Como en los test las funciones las necesitariasmos testiar para tener una mayor cobertura
// vamos a convertir esta funcion en un tipo de variable:
/*
func GetPersonByDNI(dni string) (Person, error) {
	time.Sleep(5 * time.Second) // simulamos lectura a DB
	// SELECT * FROM Persona Where ..
	return Person{}, nil
}
*/
var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second) // simulamos lectura a DB
	// SELECT * FROM Persona Where ..
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{}, nil

}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	e, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}

	// si no hay errores vamos a recibir e
	ftEmployee.Employee = e

	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Person = p

	return ftEmployee, nil

}
