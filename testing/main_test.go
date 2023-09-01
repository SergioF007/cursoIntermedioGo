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
		{25, 26, 51},
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

func TestFib(t *testing.T) {

	tables := []struct {
		a int
		b int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tables {

		fib := Fibonacci(item.a)

		if fib != item.b {
			t.Errorf("Fibonacci was incorrect, got %d expected %d", fib, item.b)
		}
	}
}
