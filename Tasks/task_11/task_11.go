// Реализовать пересечение двух неупорядоченных множеств.

package main

import (
	"fmt"
	"reflect"
	"sync"
)

type ConcurrencySet struct {
	Set   []any
	Mutex sync.RWMutex
}

func Intersection(setOne, setTwo []any) []any {
	if reflect.TypeOf(setTwo) != reflect.TypeOf(setOne) {
		return nil
	}

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

func IntersectionWithConcurrency(setOne, setTwo []any) *ConcurrencySet {
	if reflect.TypeOf(setTwo) != reflect.TypeOf(setOne) {
		return nil
	}
	set := ConcurrencySet{Set: make([]any, 0, min(len(setOne), len(setTwo)))}

	wg := sync.WaitGroup{}

	for index, i := range setOne {
		wg.Add(1)
		go func(index int, i any) {
			defer wg.Done()
			for _, j := range setTwo {
				if i == j {
					set.Mutex.Lock()
					set.Set = append(set.Set, i)
					set.Mutex.Unlock()
					return
				}
			}
		}(index, i)
	}
	wg.Wait()
	return &set
}

func main() {
	setOneInt := []any{1, 5, 2, 4, 11, 23, 6}
	setTwoInt := []any{1, 5, 7, 42, 121, 23, 26}

	res := Intersection(setOneInt, setTwoInt)
	fmt.Println(res)
	resCon := IntersectionWithConcurrency(setOneInt, setTwoInt)
	fmt.Println(resCon.Set)

	setOneFloat := []any{1.2, 5, 2.2, 4, 11, 23.1, 6}
	setTwoFloat := []any{1, 5.3, 7, 2.2, 121, 23, 26}

	res = Intersection(setOneFloat, setTwoFloat)
	fmt.Println(res)
	resCon = IntersectionWithConcurrency(setOneFloat, setTwoFloat)
	fmt.Println(resCon.Set)
}
