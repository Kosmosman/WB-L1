package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type ConcurrencyMap struct {
	Map   map[int]int
	Mutex sync.RWMutex
}

type KeyValue struct {
	Key   int
	Value int
}

// WriterIntInt Добавляем случайное значение пары ключ-значение в map (оба имеют тип int)
func WriterIntInt(cm *ConcurrencyMap, dataStream chan KeyValue, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		i, ok := <-dataStream
		if !ok {
			return
		}
		cm.Mutex.Lock()
		cm.Map[i.Key] = i.Value
		fmt.Printf("Key-value %d-%d added in %d goroutine\n", i.Key, i.Value, id)
		cm.Mutex.Unlock()
	}
}

// DataCreator Добавляем 100 случайных значений в канал, после чего закрываем его
func DataCreator(dataStream chan KeyValue, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(dataStream)

	for i := 0; i < 100; i++ {
		tmp := KeyValue{Key: rand.Int(), Value: rand.Int()}
		dataStream <- tmp
	}
}

func main() {
	cm := ConcurrencyMap{Map: make(map[int]int)}
	wg := sync.WaitGroup{}
	dataStream := make(chan KeyValue)

	wg.Add(1)
	go DataCreator(dataStream, &wg)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go WriterIntInt(&cm, dataStream, i, &wg)
	}
	wg.Wait()
}
