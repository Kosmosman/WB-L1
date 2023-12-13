// Реализовать пересечение двух неупорядоченных множеств.

package main

import (
	"fmt"
)

func Intersection(setOne, setTwo []any) []any {
	var result = make([]interface{}, 0, min(len(setOne), len(setTwo)))

	for _, i := range setOne {
		for _, j := range setTwo {
			if i == j {
				result = append(result, i)
				break
			}
		}
	}
	return result
}

func main() {
	setOneInt := []any{1, 5, 2, 4, 11, 23, 6}
	setTwoInt := []any{1, 5, 7, 42, 121, 23, 26}
	res := Intersection(setOneInt, setTwoInt)
	fmt.Println(res)
	setOneInt = []any{1.2, 5, 2.2, 4, 11, 23.1, 6}
	setTwoInt = []any{1, 5.3, 7, 2.2, 121, 23, 26}
	res = Intersection(setOneInt, setTwoInt)
	fmt.Println(res)
}
