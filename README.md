# cursoIntermedioGo
Go orientado a objetos - onceptos de POO en Golang -  Test Unitarios en Go para aplicar TDD - Calcular el Code Coverage de mi proyecto  - Análisis del Profiling en tus programas - Cómo multiplexar mensajes de canales - Técnicas para la creación de Worker Pools concurrentes - Crear un Job Queue concurrente

# Curso de Go Intermedio: Programación Orientada a Objetos y Concurrencia

## ****Características esenciales de Go****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/edab3244-01cb-42e5-a56b-b524df5a56e1/Untitled.png)

- **Go es Fuertemente Tipado   (adicional)**

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/8a132490-f2a9-45fc-9146-56dd3eb5bd8d/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/f9726b6b-bbeb-4da7-a18d-0517c41b1220/Untitled.png)

**Conocimientos Requeridos**

- Declaración de variables.
- Condicionales.
- Sintaxis básica.
- Declaración de GoRoutines y Channels.
- Slices y maps.
- Apuntadores.

**¿Qué terminarás sabiendo?**

- ¿Es Go orientado a objetos?
- Cómo aplicar los conceptos de POO en Golang.
- Crear Test Unitarios en Go para aplicar TDD.
- Calcular el Code Coverage de mi proyecto.
- Análisis del Profiling en tus programas.
- Cómo multiplexar mensajes de canales.
- Técnicas para la creación de Worker Pools concurrentes.
- Crear Test Unitarios en Go para aplicar TDD.
- Crear un Job Queue concurrente.

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/41f9d930-49b2-4fd2-b7d9-36c1405dbe87/Untitled.png)

## ****Repaso general: GoRoutines y apuntadores****

```go

func doSomething() {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
```

Por lo que en main, primero tenemos que **crear el canal**  y luego si definir la funcion con **Go**

```go
c := make(chan int) 
go doSomething(c)
<-c
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/0c829912-bc30-4a18-bb63-57e021d222a9/Untitled.png)

### Apuntadores:

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/b5f20606-9749-49df-acdf-3fe0eeb617f4/Untitled.png)

## ****¿Es Go orientado a objetos?****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/81fabd34-af07-4f3b-886f-d29abb5c48d1/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/6eeb104d-2cef-4a4c-8444-6e874265fd6e/Untitled.png)

## ****Structs vs. clases****

para poder ejecutar cambión en ubuntu wsl tengo que otorgar los permisos correspondientes:

 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/798cb5b9-40ff-4517-9dc0-1dcf4c5f8d2a/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/20aa74e5-bfd7-421b-a778-7c4ed1d89876/Untitled.png)

## ****Constructores****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/82a6392a-4079-4e41-bc29-4a6cc2c95fc3/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/b4a7f97e-3ef6-44b0-9ba4-d53cc8a2f964/Untitled.png)

### Vamos a ver una forma de emular el comportamiento de constructores mediante una función.

Lo que hacemos es crear una función a la cual la llamaremos NewEmployee y como parametros de entradas definimos los parámetros con el cual queramos construir el objeto y el return lo redirigimos con el puntero a Employee. De la siguiente forma: 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/befc921e-5c27-467d-8ae7-8b730439cf16/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/de2ed675-1ad8-48df-ae5d-a3cab0bbef31/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/41b19c9f-4f1b-4dda-a905-f6dbbb7519c6/Untitled.png)

## ****Herencia****

```go
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
	GetMessage(ftEmployee) // Por lo que este tipo de acciones las podemos representar por medio de interfaces

}
```

### Nota importante:

**Me genera erro por que go no permite implemetar herencia, por lo tanto.
puedo aplicarla hasta cierto punto como es el caso de los atributos en en los métodos el comportamiento no se hereda**
	GetMessage(ftEmployee) **// Por lo que este tipo de acciones las podemos representar por medio de interfaces**

## ****Interfaces****

son como contratos los cuales si se implemeta se tiene que comprometer a aplicar todos sus metodos. 

```go
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
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/481c0ef4-7523-4b3a-91bc-7e58b1efcf0d/Untitled.png)

## ****Aplicando interfaces con Abstract Factory****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/efa1d2eb-d94d-4563-a084-8ff571e1c907/Untitled.png)

Problema: Crear un diseño de software que sea capaz de enviar notificaciones tanto de tipo SMS como de Email  utilizando **Polimorfismo e Interfaces**

```go
package main

import "fmt"

// Factoru para el SMS

type INotificationFactory interface {
	SendNotification() // Enviar notificacion
	GetSender() ISender // Obtener remitente
}

type ISender interface {
	GetSenderMethod() string  // Obtener método del remitente
	GetSenderChannel() string  // Obtener canal del remitente
}

type SMSNotification struct {

}

// estamos creando el metodo para SendNotification
func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS")
}

func (SMSNotification) GetSender() ISender{
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {

}

func (SMSNotificationSender) GetSenderMethod() string{
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string{
	return "Twilio"
}
```

## ****Implementación final de Abstract Factory****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/443878f2-defc-4d71-ace8-4ae0478bc825/Untitled.png)

```go
package main

import "fmt"

// Factory para el SMS

type INotificationFactory interface {
	SendNotification()  // Enviar notificacion
	GetSender() ISender // Obtener remitente
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

// estamos creando el metodo para SendNotification
func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}  // obtengo la notificacion del remitente
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// Factory para el Email

type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notification via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	//suponemos que estamos implemetnado un servicio de amazon para enviar el correo.
	return "SES"
}

// Creamos la funcion para que nos retorne las clases concretas que quermos utilizar para enviar las notificaciones
// resivimos un tipo de notificacion como parametro (para saber cual es el que se tiene que enviar )
// Vamos a retornar es un INotificationFactory o  un error
func getNotificationFactory(notificationType string) (INotificationFactory, error) {

	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("Notification Type")
}

// Creao la funcion que me captura cual es la notificacion que se va enviar
// Enviar notificación
func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {

	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/f654fd0e-7dd1-4c96-9fe9-b6345db52bb3/Untitled.png)

## ****Funciones anónimas****

El primer ejemplo son funciones anónimas que se ejecutar en el ****Goroutin del main****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/fdababe3-917d-444c-af71-5d8825b7c9b6/Untitled.png)

En este ejemplo es una funcione anónima que se ejecuta al mismo tiempo que el **Goroutin del main, al crear un canal y asignarselo a la función que se pretende ejecutar de manera concurrente.** 

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	go func() {
		fmt.Println("Starting Function")
		time.Sleep(5 * time.Second)
		fmt.Println("End")
		c <- 1
	}()

	<-c
	close(c)
}
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/805db118-4f06-4c1c-86a7-e706b8dd9072/Untitled.png)

## ****Funciones variadicas y retornos con nombre****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/3305f014-96db-4aaf-9831-4acfebf17518/Untitled.png)

```go
package main

import "fmt"

func sum1(a, b int) int {
	return a + b
}

func sum3(a, b, c int) int {
	return a + b + c
}

// Para sumar N numeros
// el argumento values es tratado com Slice
func sum(values ...int) int {
	total := 0
	for _, num := range values {

		total += num
	}
	return total
}

// con String
func printName(values ...string) {
	for _, name := range values {
		fmt.Println(name)
	}
}

// devolver datos multiples con una solo entrada
func getValues(x int) (int, int, int) {
	return 2 * x, 3 * x, 4 * x
}

// ó

func getValues2(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(sum(1, 2, 3))
	printName("Sergio", "Paula", "Kike") // dentro de la funcion se ejecuta el metodo fmt. 
	fmt.Println(getValues(10))
	fmt.Println(getValues2(20))
}
```

## ****Cómo utilizar los Go modules****

Creamos un paquete nuevo llamado ****************modules**************** y desde la terminal creamos el archivo .mod de la siguiente forma: 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/ce507d51-c3cb-4de7-89bd-c085682238e3/Untitled.png)

Descargamos un librería con el siguiente comando:

 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/1e351476-4502-4fb2-ac91-59f14131706c/Untitled.png)

y vemos como nuestro archivo go.mod se nos actualiza con esta nueva librería y se crea el archivo go.sum. 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/56c30748-49b3-41a6-a04d-e78824c2cc66/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/79dbd8c9-1ace-43c5-904d-dea3b490884f/Untitled.png)

en el caso que quieras eliminar las dependencias de la librerías que ya no estan en uso usa el siguiente comando.  

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/4cba7064-71e2-4853-b31d-2bd9c95eb13f/Untitled.png)

## Nota: go mod tidy

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/0f6ca583-091c-4971-8bf2-a5eb070af852/Untitled.png)

Ej: utilización de una librería 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/4a34b3be-dc3a-4e06-bbef-51e948db6423/Untitled.png)

Descargamos una versión 2 de esa librería y vemos como implementamos, para la versión 2 toca definirle un ******Alias****** para que no genere error y que go las pueda identificar. 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/f8d42670-e104-4038-8ecd-03d30738c601/Untitled.png)

## ****Creando nuestro módulo****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/25b988ea-bc57-4738-9cd7-95f1ae90b145/Untitled.png)

Creamos un paquete que implementa un println,  donde creamos el mod en la linea de comandos: 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/19badd82-1a56-4e22-9a71-10c612f1bc24/Untitled.png)

```go
package utils

import "fmt"

// funcion publica unicia en Mayuscula
func HelloWorld() {
	fmt.Println("Hello World From Platzi")
}
```

Luego creamos un el git init de este paquete para cargarlo a un repo de github 

creamos otro paquete donde vamos a implementar el modulo que acabamos de crear. 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/a3e7408f-7461-4a98-935e-19613f9362ea/Untitled.png)

vemos que se nos instala la libreria que cargamos en githut

```go
package main

import utils "github.com/SergioF007/HELLOPLATZIMOD"

func main() {

	utils.HelloWorld()

}
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/d9d026e3-36d3-4cb3-9a31-edeabef19856/Untitled.png)

## ****Testing****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/48a7e012-2103-418d-ae18-65aa1a9acf33/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/96b64320-049c-43be-b0e9-11ebdc248f19/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/51cc8bea-1043-41f6-89b3-4dff9f0e8b79/Untitled.png)

Generamos un test para la función Sum  y en el test creemos un arreglo donde un objeto llamado tables, donde **a** y **b** son los atributos que se van a sumar y **n** es el resultado de la **suma.** Creamos un for para capturar los objetos de tables y como nos interesa es su contenido capturamos los items. 

## ****Code coverage****

Para identificar la cobertura de nuestro test 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/3f296b95-6103-44a3-81ce-33cb8ad7eeba/Untitled.png)

**Para ver la cobertura mas explícitamente** 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/ca37101e-bca8-40ce-955f-1e3fb973e217/Untitled.png)

**Y para poderlo ver mas visualmente en un archivo Html**

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e9e98d7d-c018-470f-ab0a-a40743ecd5c7/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/8e64e06c-bdab-4777-ab82-79ca77684733/Untitled.png)

```go
package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	// total := Sum(5, 5)

	// if total != 10 {
	// 	t.Errorf("Sum was incorrect, go %d expected %d", total, 10)
	// }

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
	
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		//validamos su la suma de a y b es igual a n
		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.n)
		}
	}

}

func TestMax(t *testing.T) {

	tables := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{10, 5, 10},
		{10, 15, 15},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)

		if max != item.n {
			fmt.Printf("GetMax was incorrect, got %d, expected %d", max, item.n)
		}
	}

}
```

Implemente el tests tambien para función GetMazx y evalue de nuevo la cobertura del test. 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/ef65579b-c986-480f-bcd2-646fbd05d2ad/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/5bf92a25-45b6-40cc-ab69-df40599c6121/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/bf9aa539-32e6-4f44-ae6e-5ac56c031335/Untitled.png)

## ****Profiling****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/7a1825ba-efad-4387-8739-a2c9f6c662b0/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/aacc029a-6dec-45ec-a4df-e31de9b009e7/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/5927ef48-c14d-4d02-b4af-2b4fe226eeb2/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/8c782abf-83a0-43e3-a195-d505c3c8b57b/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/de8543cc-5f78-4dda-9231-03ed476f408e/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/3a7bb6dd-8b60-411f-ba23-76ffd416122d/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/864f0973-9742-4f5b-83ee-c89f044f9480/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/c3cb00d6-fa1c-436e-ba82-4bd2d61133bb/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/9a8dabf9-3fd3-498b-a44e-9d6637907269/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e1201902-a7d0-4259-b13c-c40189bdd70d/Untitled.png)

## ****Testing usando Mocks****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/96d106f5-4af5-4fff-988c-d31d3188ad5c/Untitled.png)

**************main.go**************

```
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
```

************************main_test.go************************

```go
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
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e1711857-6567-44e4-9942-95c376624c7c/Untitled.png)

## ****Unbuffered channels y buffered channels****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/d001d73a-6a8a-47f0-8a2f-250b415a3b96/Untitled.png)

```go
package main

import "fmt"

func main() {

	// canal sin buffer: c := make(chan int, 1)
	// canal con buffer
	c := make(chan int, 1)  // canal para 1 elemento
	ch := make(chan int, 3) // canal para 3 elementos

	ch <- 1
	ch <- 2
	ch <- 3

	c <- 1 // enviamos un valor por ese canal

	fmt.Println(<-c)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/a20e784f-d645-483d-b64d-cbc3e1132bf8/Untitled.png)

Los canales son necesarios para poder comunicar nuestras rutinas, ya que si no aplicamos canales, terminará la ejecución antes de que todas las rutinas las completen.  

## ****Waitgroup****

### Nota: Recuerda que si no tenemos canales lo que pasara es que van a terminar su ejecución antes de que las demás rutinas termine su ejecución. La otra alternativa para tener este tipo de sincronización entre la rutinas es el concepto de ****Waitgroup****

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {

		wg.Add(1)
		go doSomething(i, &wg)

	}

}

func doSomething(i int, wg *sync.WaitGroup) {
	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("End")
}
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/7c83f9f4-3cca-488b-9d37-1501e7939618/Untitled.png)

EL ****Waitgroup es un contador que se utiliza para saber cuando se lanza una goroutines sumando 1 con el   wg.Add(1)   luego en la función cuando se terminar de ejecutar lo pone en cero con****  wg.Done()  y para garantizar que el que todas las ****goroutines terminaron usamos el**** wg.Wait() que me identifica cuando el contador esta en cero. 

### Nota: Esto nos permitira ejecutar nuestro programa sin necesidad de utilizar un canal.

****************************El Ejemplo con canales seria de la siguiente manera:**************************** 

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 10)

	for i := 1; i <= 10; i++ {
		wg.Add(1) // Sumo cada ves que se lance una rutina 
		go doSomething(i, c)

	}

	for i := 1; i <= 10; i++ {
		<-c // Esperamos a que todas las goroutines terminen
	}

	wg.Wait() // cuando este contador llegue a 0 el programa va estar bloqueado hasta que pase este proceso

}

// hacer algo
func doSomething(i int, c chan int) {
	defer wg.Done() // cuando se fianalice la rutina pongo el 0 el contador
	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("End")
	c <- i

}
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/dc9a597f-0f9d-495a-959a-8c7bfd4fc8ed/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/2cc0c63a-242a-4b65-9498-e04287b29542/Untitled.png)

## ****Buffered channels como semáforos****

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	c := make(chan int, 2)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {

		c <- 1 // el canal esta actuando como semaforo, ya que por cada rutina ejecutamos 2 grountien al tiempo
		wg.Add(1)
		go doSomething(i, &wg, c)

	}

	wg.Wait()

}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Id %d started\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("Id %d finished\n", i)
	<-c
}
```

******************************************Modificando el Buffer del canal puede indicar cuantas gourtens se va a ajecutar al tiempo****************************************** 

Por lo tanto este comportamiento empieza a actuar como semáforo ! 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/38354628-a069-4e49-ad50-46f8152ee0f0/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/54d2f3e8-a8f5-45e9-a5fb-b5d2a263a9f3/Untitled.png)

## ****Definiendo channels de lectura y escritura****

Como buena practica, al momento de utilizar canales es recomendables definirles de que tipo va hacer, si solo de Escritura o solo de lectura o Ambas, Como el **siguiente Ejemplo:** 

```go
package main

import "fmt"

// Solo de escritura por eso: chan<-
func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

// va a tener un canar de entrada y uno de salida
// el primer canar es exclusivamente de lectrua(in <-chan int) y el segundo canar es exclusivamente de escritura(out chan<- int)
func Double(in <-chan int, out chan<- int) {
	for value := range in {
		// devolvemos el dato del canal de entrada al doble por el canal de salida
		out <- 2 * value
	}
	close(out)
}

// funcion de imprimir por cada canal de tipo int
// exclusivo de tipo lectura <-chan
func Print(c <-chan int) {

	for value := range c {
		fmt.Println(value)
	}
}

func main() {

	// vamos a crear los canales que se van a implememtar en las funciones

	generetor := make(chan int)
	doubles := make(chan int)

	go Generator(generetor)
	go Double(generetor, doubles)

	Print(doubles) // lemos el canal se salida de Double

}
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/8dcd9a9b-ea35-406c-b485-304726487d8c/Untitled.png)

## ****Worker pools****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/16aad439-71da-4a85-a5c4-d20078e8f8e4/Untitled.png)

```go
package main

import "fmt"

func Worker(id int, jobs <-chan int, results chan<- int) {

	for job := range jobs {

		fmt.Printf("Worker with id %d started fib with %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d, job %d and fib %d\n", id, job, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {

	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {

	tasks := []int{1, 2, 3, 4, 5, 6, 10, 22, 40}
	nWorkers := 3

	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {

		go Worker(i, jobs, results)

	}

	for _, value := range tasks {
		jobs <- value
	}

	close(jobs)

	// vamos a leer los resultados

	for i := 0; i < len(tasks); i++ {
		<-results // leemos los resultados
	}

}
```

******Probar:****** 

Escribí unas modificaciones del código visto en clase para observar realmente *cómo es que las gorutines y los workers funcionan a gran escala* (o bueno eso creo), el código al final de su ejecución dice cuánto tiempo tomó y en mi caso, realizando las series de Fibonacci con los workers y gorutines tomó **9.544643109** segundos, de la forma convencional tomó **18.054724068** les dejo el código por si alguien quiere probarlo, sólo deben cambiar el valor de la constante withGoRutinesAndWorkers para hacer la ejecución de una forma u otra.

**Código:**

```go
package main

import (
	"fmt"
	"time"
)

funcWorker(id int, jobs <-chanint, resultschan<-int) {
	for job :=range jobs {
			fmt.Printf("Worker with id %d started fib with %d\n", id, job)
			fib := Fibonacci(job)
			fmt.Printf("Worker with id %d has finished the job %d, and fib %d\n", id,job,fib)
			results <- fib
		}
}

funcFibonacci(nint)int {
if n <= 1 {
return n
	}
return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	start := time.Now()
	tasks := []int{2, 34, 36, 17, 10, 40, 6, 22, 41, 44, 33, 2, 5, 2, 34, 36, 17, 10, 40,   6, 22, 41, 44, 33, 2, 5, 17, 10, 40, 6, 22, 2, 5, 2, 34, 36, 22, 41, 44}
const withGoRutinesAndWorkers =falseif withGoRutinesAndWorkers {
		nWorkers := 12
		jobs := make(chanint, len(tasks))
		results := make(chanint, len(tasks))

for i := 0; i <= nWorkers; i++ {
go Worker(i, jobs, results)
		}

for _, value :=range tasks {
			jobs <- value
		}
		close(jobs)

for i := 0; i < len(tasks); i++ {
			<-results
		}
	}else {

for _, v :=range tasks {
			val := Fibonacci(v)
			fmt.Println("val->", val)
		}

	}

	elapsed := time.Since(start)
	fmt.Printf("Program took %s", elapsed)

}
```

## ****Multiplexación con Select y Case****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/870031cb-cb5c-4214-8a8e-db57f81efc8a/Untitled.png)

El siguiente ejemplo vamos hacer que un canal tarde mas que otro para ver cual termina primero: 

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int)
	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go DoSomething(d1, c1, 1)
	go DoSomething(d2, c2, 2)

	fmt.Println(<-c1)
	fmt.Println(<-c2)
}

func DoSomething(i time.Duration, c chan<- int, param int) {

	time.Sleep(i)
	c <- param

}
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/0c4b65ad-bb39-4287-b99e-a12fd9bbd749/Untitled.png)

Pero vemos que esto no paso, por lo que primero imprime param del canal 1 siendo al que mas se le definio el tiempo de espera, esto ocurre por que el programa sabe cual es el que termina primero por que no tenemos algo que nos capture esta respuesta. Asi que para eso implementamos el select: 

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int)
	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go DoSomething(d1, c1, 1)
	go DoSomething(d2, c2, 2)

	for i := 0; i < 2; i++ {

		select {
		case channelMsg1 := <-c1:
			fmt.Println(channelMsg1)
		case channelMsg2 := <-c2:
			fmt.Println(channelMsg2)
		}

	}
}

func DoSomething(i time.Duration, c chan<- int, param int) {

	time.Sleep(i)
	c <- param

}
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/f39cfbc8-ef73-49a0-ae61-d1667c794211/Untitled.png)

## Vamos a crear un proyecto (Web Services)  que va hacer capaz de generar Jobs que un ****Worker pools va a estar consumiendo****

### ****Definiendo workers, jobs y dispatchers****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/dc843c54-34d8-41c0-a1a0-b9df384a8c23/Untitled.png)

Un pequeño resumen de lo que pude entender, espero los ayude

- El dispatcher recibe todos los jobs, se puede decir que es como el componente global
- Cada worker tiene su canal de jobs, y saben cual es el canal del disptacher, es decir el workerpool es el mismo canal para todos los workers.
- Cada worker esta enviando su canal al canal del dispatcher
- En la medida que el dispatcher recibe jobs este los va repartiendo entre los workers a través de sus canales

!https://i.ibb.co/DGtYx9G/Untitled.png

```go
package main

import (
	"fmt"
	"time"
)

type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

type Worker struct {
	Id         int
	JobQueue   chan Job      // JobQueue (Cola de trabajos)
	WorkerPool chan chan Job // es un canal de canales
	QuitChan   chan bool
}

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

func NewWorker(id int, workerPool chan chan Job) *Worker {

	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool),
	}

}

func (w Worker) Start() {

	go func() {
		for {
			w.WorkerPool <- w.JobQueue

			select {
			case job := <-w.JobQueue:
				fmt.Printf("Worker with id %d Started\n", w.Id)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker with id %d Finished with result %d \n", w.Id, fib)
			case <-w.QuitChan: // para cuando reciba que tiene que cerrarce el canal.
				fmt.Printf("Worker with id %d Stopped \n", w.Id)
			}
		}
	}()
}

func (w Worker) Stop() {

	go func() {
		w.QuitChan <- true
	}()
}

func Fibonacci(n int) int {

	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {

	worker := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: worker,
	}
}

// cremos el metodo para que encole los Jobs
func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:

			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job

			}()
		}
	}
}
```

## ****Creando web server para procesar jobs****

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

type Worker struct {
	Id         int
	JobQueue   chan Job      // JobQueue (Cola de trabajos)
	WorkerPool chan chan Job // es un canal de canales
	QuitChan   chan bool
}

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

func NewWorker(id int, workerPool chan chan Job) *Worker {

	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool),
	}

}

func (w Worker) Start() {

	go func() {
		for {
			w.WorkerPool <- w.JobQueue

			select {
			case job := <-w.JobQueue:
				fmt.Printf("Worker with id %d Started\n", w.Id)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker with id %d Finished with result %d \n", w.Id, fib)
			case <-w.QuitChan: // para cuando reciba que tiene que cerrarce el canal.
				fmt.Printf("Worker with id %d Stopped \n", w.Id)
			}
		}
	}()
}

func (w Worker) Stop() {

	go func() {
		w.QuitChan <- true
	}()
}

func Fibonacci(n int) int {

	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {

	worker := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: worker,
	}
}

// cremos el metodo para que encole los Jobs
func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:

			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job

			}()
		}
	}
}

func (d *Dispatcher) Run() {

	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start() //  les doamos inicio
	}

	go d.Dispatch()

}

// esta funcion es la encargada de manejar todo lo que venga en las peticiones HTTP del Servidor
func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {

	if r.Method != "POST" { // GET, PUT, DELETE

		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed) // esto le respondera al cliente que no podra utilizar este metodo

	}

	//FormValue nos permite acceder a todo lo que se esta enviando en el Request y de esta forma son capaz de extraerlo
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Invalid Delay", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid Value", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "Ivalid Name", http.StatusBadRequest)
		return
	}

	job := Job{Name: name, Delay: delay, Number: value}
	//Lo encolamos con el canal de jobQueue
	jobQueue <- job
	w.WriteHeader(http.StatusCreated)

}

func main() {

	const (
		maxWorkers   = 4       // Definimos la cantidad de Worker con los que vamos a trabajar
		maxQueueSize = 20      // Definimos la maxima cantidad de Jobs que van a ser procesados simultaneamente
		port         = ":8081" // definimos por cual puerto nuestro servidor va a estar escuchando
	)

	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)

	dispatcher.Run()
	// http://localhost:8081/fib
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})
	log.Fatal(http.ListenAndServe(port, nil))

}
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/1c9fbc11-5957-428d-9844-e3b424a69993/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/da873434-01a0-4fd3-8ab5-042b0e42abbb/Untitled.png)

### Resumen

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/81415504-480c-4ac1-8eaf-7bd11ed52d80/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/d12af821-c00a-4ba3-9b4b-cabcfb9ec1bc/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/8558bee5-1fad-4d27-9a60-32ce146a0c1a/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/37b297b7-270b-410c-8b88-cc458a98e31c/Untitled.png)
