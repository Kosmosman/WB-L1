// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"fmt"
	"math/rand"
)

// Быстрая сортировка с использованием слияния, вследствие чего создается новый отсортированный массив
func Sort(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	if len(arr) <= 2 {
		if arr[0] > arr[len(arr)-1] {
			arr[0], arr[len(arr)-1] = arr[len(arr)-1], arr[0]
		}
		return arr
	}
	pivot := arr[rand.Intn(len(arr))]
	equalValues := make([]int, 0, len(arr))
	lowerValues := make([]int, 0, len(arr))
	highestValues := make([]int, 0, len(arr))
	for _, v := range arr {
		if v == pivot {
			equalValues = append(equalValues, v)
		} else if v > pivot {
			highestValues = append(highestValues, v)
		} else {
			lowerValues = append(lowerValues, v)
		}
	}
	return append(append(Sort(lowerValues), equalValues...), Sort(highestValues)...)
}

// Быстрая сортировка с перемещением элементов на исходном массиве, создание дополнительным массивов не требуется
func InPlaceSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	pivot := arr[rand.Intn(len(arr))]
	var less, equal, higher int
	for higher = 0; higher < len(arr); higher++ {
		if arr[higher] < pivot {
			arr[higher], arr[less] = arr[less], arr[higher]
			if arr[less] != arr[equal] {
				arr[higher], arr[equal] = arr[equal], arr[higher]
			}
			less++
			equal++
		} else if arr[higher] == pivot {
			arr[higher], arr[equal] = arr[equal], arr[higher]
			equal++
		}
	}
	InPlaceSort(arr[:less])
	InPlaceSort(arr[equal:])
}

func main() {
	example := []int{2, 6, 1, 5, 2, 8, 3, 2, 5, 9, 20, -1, 7}
	fmt.Printf("Before sorting        %v\n", example)

	res := Sort(example)
	fmt.Printf("After sort with merge %v\n", res)

	InPlaceSort(example)
	fmt.Printf("After inplace sort    %v\n", example)
}
