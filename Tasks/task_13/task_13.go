package main

import (
	"fmt"
	"log"
)

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Before: a = %d, b = %d\n", a, b)

	// Обмен числами
	if a != b {
		a ^= b
		b ^= a
		a ^= b
	}

	fmt.Printf("After: a = %d, b = %d", a, b)
}
