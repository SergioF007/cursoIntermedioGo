package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	c := make(chan int, 10)
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
