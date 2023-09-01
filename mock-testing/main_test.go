package main

import (
	"fmt"
	"testing"
)

func TestGetFullTimeEmployeeById(t *testing.T) {

	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}
				GetPersonByDNI = func(id string) (Person, error) {
					return Person{
						Name: "Sergio Fernanadez",
						Age:  27,
						DNI:  "1",
					}, nil
				}
			},

			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Name: "Sergio Fernanadez",
					Age:  27,
					DNI:  "1",
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}
	// implementamos las siguientes lineas, para evitar que se siga haciendo de manera permanente
	originalGetEmpliyeeById := GetEmployeeById
	originalGetPersonByDni := GetPersonByDNI
	// por eso al final debemos ponerlos en su estado originar.

	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeById(test.id, test.dni)

		if err != nil {
			fmt.Printf("Error when getting Employee")
		}

		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error, got %d expected %d", ft.Age, test.expectedEmployee.Age)
		}

	}

	GetEmployeeById = originalGetEmpliyeeById
	GetPersonByDNI = originalGetPersonByDni
}
