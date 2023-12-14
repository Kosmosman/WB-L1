// Удалить i-ый элемент из слайса.
package main

import (
	"errors"
	"fmt"
	"log"
)

// Удаление элемента с сохранением порядка остальных элементов
func DeleterWithSaveOrder(arr []any, deletedElem int) ([]any, error) {
	if deletedElem < 0 || deletedElem >= len(arr) {
		return []any{}, errors.New("incorrect index")
	}
	if deletedElem == 0 {
		return arr[1:], nil
	}
	if deletedElem == len(arr)-1 {
		return arr[:len(arr)-1], nil
	}
	copy(arr[deletedElem:], arr[deletedElem+1:])
	return arr[:len(arr)-1], nil
}

// Удаление элемента без сохранения подрядка остальных, более эффективно, чем с сохранением
func DeleterWithoutSaveOrder(arr []any, deletedElem int) ([]any, error) {
	if deletedElem < 0 || deletedElem >= len(arr) {
		return []any{}, errors.New("incorrect index")
	}
	arr[len(arr)-1], arr[deletedElem] = arr[deletedElem], arr[len(arr)-1]
	return arr[:len(arr)-1], nil
}

func TestDeleteWithSave() {
	arr := []any{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Original array : \t\t\t   %v\n", arr)
	for i := 0; i < len(arr); i++ {
		delArr, err := DeleterWithSaveOrder(arr, i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Array after deleted element with index %d : %v\n", i, delArr)
		arr = []any{0, 1, 2, 3, 4, 5, 6, 7, 8}
	}
}

func TestDeleteWithoutSave() {
	arr := []any{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Original array : \t\t\t   %v\n", arr)
	for i := 0; i < len(arr); i++ {
		delArr, err := DeleterWithoutSaveOrder(arr, i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Array after deleted element with index %d : %v\n", i, delArr)
		arr = []any{0, 1, 2, 3, 4, 5, 6, 7, 8}
	}
}

func main() {
	TestDeleteWithSave()
	println()
	TestDeleteWithoutSave()
}
