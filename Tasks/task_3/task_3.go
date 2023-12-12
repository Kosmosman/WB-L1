//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов с использованием конкурентных вычислений

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
func SumSquaresWithAtomic() {
	var pos atomic.Int32
	var sum atomic.Int32
	wg := sync.WaitGroup{}
	arr := GenerateArray(20)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(arr *[]int, pos *atomic.Int32, sum *atomic.Int32, wg *sync.WaitGroup) {
			defer wg.Done()

			posValue := pos.Load()
			for int32(len(*arr)) > posValue {
				pos.Add(1)
				sum.Add(int32((*arr)[posValue] * (*arr)[posValue]))
				posValue = pos.Load()
			}
		}(&arr, &pos, &sum, &wg)
	}
	wg.Wait()
	println(sum.Load())
}

// Расчет с использованием горутин, waitGroup и mutex для синхронизации позиции расчетного числа
func SumSquaresWithMutex() {
	wg := sync.WaitGroup{}
	arr := GenerateArray(20)
	mutex := sync.Mutex{}
	var pos int
	var sum int

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(arr *[]int, pos *int, sum *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
			defer wg.Done()

			var currentPos int

			for {
				mutex.Lock() // Начало лока
				if *pos >= len(*arr) {
					mutex.Unlock()
					break
				}
				currentPos = *pos
				*pos++
				*sum += (*arr)[currentPos] * (*arr)[currentPos]
				mutex.Unlock() // Конец лока
			}
		}(&arr, &pos, &sum, &wg, &mutex)
	}
	wg.Wait()
	println(sum)
}

func main() {
	SumSquaresWithMutex()
	println("\n----------------------\n")
	SumSquaresWithAtomic()
}
