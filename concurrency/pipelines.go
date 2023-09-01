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
// el primer canar es exclusivamente de lectrua y el segundo canar es exclusivamente de escritura
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

	Print(doubles) // lemos el canal se dalida de Double

}
