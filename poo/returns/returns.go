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
	printName("Sergio", "Paula", "Kike")
	fmt.Println(getValues(10))
	fmt.Println(getValues2(20))
}
