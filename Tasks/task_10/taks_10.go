// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

package main

import (
	"fmt"
	"math"
	"sync"
)

type ConcurrencyMap struct {
	Map   map[float64][]float64
	Mutex sync.Mutex
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	wg := sync.WaitGroup{}
	m := ConcurrencyMap{Map: make(map[float64][]float64)}

	// Проходим по всему срезу, и для каждого элемента выделям по горутине, ищущей подходящие значения в группу
	for i := range temperatures {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for _, value := range temperatures {
				if math.Abs(value-temperatures[i]) <= 10 {
					m.Mutex.Lock()
					m.Map[temperatures[i]] = append(m.Map[temperatures[i]], value)
					m.Mutex.Unlock()
				}
			}
		}(i)
	}

	wg.Wait()

	for k, v := range m.Map {
		fmt.Printf("%f - ", k)
		for _, value := range v {
			fmt.Printf("%f ", value)
		}
		println()
	}
}
