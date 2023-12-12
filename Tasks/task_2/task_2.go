// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
// и выведет их квадраты в stdout.

package main

import (
	"sync"
	"sync/atomic"
)

func GenerateArray(count int) []int {
	var result = make([]int, count, count)
	for i := 0; i < count; i++ {
		result[i] = i
	}
	return result
}

// Расчет с использованием горутин, waitGroup и atomic-счетчика для определния текущей нерасчитанной позиции в слайсе
func CalculateSquaresWithAtomic() {
	var pos atomic.Int32
	wg := sync.WaitGroup{}
	arr := GenerateArray(100)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(arr *[]int, pos *atomic.Int32, wg *sync.WaitGroup) {
			defer wg.Done()

			posValue := pos.Load()
			for int32(len(*arr)) > posValue {
				pos.Add(1)
				println((*arr)[posValue] * (*arr)[posValue])
				posValue = pos.Load()
			}
		}(&arr, &pos, &wg)
	}
	wg.Wait()
}

// Расчет с использованием горутин, waitGroup и mutex для синхронизации позиции расчетного числа
func CalculateSquaresWithMutex() {
	wg := sync.WaitGroup{}
	arr := GenerateArray(100)
	mutex := sync.Mutex{}
	var pos int
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(arr *[]int, pos *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
			defer wg.Done()

			var currentPos int

			for {
				mutex.Lock() // Начало лока
				if *pos >= len(*arr) {
					mutex.Unlock()
					return
				}
				currentPos = *pos
				*pos++
				mutex.Unlock() // Конец лока
				println((*arr)[currentPos] * (*arr)[currentPos])
			}
		}(&arr, &pos, &wg, &mutex)
	}
	wg.Wait()
}

func main() {
	CalculateSquaresWithAtomic()
	println("\n----------------------\n")
	CalculateSquaresWithMutex()
}
