package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		c := make(chan int, 10)

		for i := 1; i <= 10; i++ {

			c <- i
			go doSomething(i, c)

		}

		for i := 1; i <= 10; i++ {
			<-c // Esperamos a que todas las goroutines terminen
		}
	*/

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {

		wg.Add(1)
		go doSomething(i, &wg)

	}

	wg.Wait()

}

/*
// hacer algo
func doSomething(i int, c chan int) {

	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("End")
	<-c

}
*/

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("End")
}
