// Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"fmt"
	"sort"
)

// В случае нахождения - возвращается индекс искомой величины, в ином случае - возвращается -1
func BinSearch(target int, arr []int) int {
	// Бинпоиск проводится только на отсортированном массиве, поэтому если подается неотсортированный массив,
	// необходимо подключить нижестоящую строчку. Однако для программ, в которых бинпоиск выполняется больше одного раза,
	// сортировка должна исполняться до начала процедуры поиска
	//sort.Ints(arr)

	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)/2 // Для защиты от переполнения
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 2, 7, 8, 4, 1, 6, 11}
	sort.Ints(arr)
	find := []int{4, 60, 2, 1}
	fmt.Printf("Array: %v\n", arr)
	for _, v := range find {
		if i := BinSearch(v, arr); i != -1 {
			fmt.Printf("Value %d founded in array on index %d\n", v, i)
		} else {
			fmt.Printf("Value %d not founded in array\n", v)
		}
	}
}
