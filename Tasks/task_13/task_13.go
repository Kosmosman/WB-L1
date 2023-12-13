// Поменять местами два числа без создания временной переменной.

package main

import (
	"fmt"
	"log"
)

func SwapWithXor(a, b int) (int, int) {
	if a != b {
		a ^= b
		b ^= a
		a ^= b
	}
	return a, b
}

func SwapWithAssignment(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Before: a = %d, b = %d\n", a, b)

	a, b = SwapWithXor(a, b)
	fmt.Printf("After xor: a = %d, b = %d\n", a, b)

	a, b = SwapWithAssignment(a, b)
	fmt.Printf("After assigment: a = %d, b = %d", a, b)
}
