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
