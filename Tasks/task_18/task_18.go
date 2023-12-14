// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

type ConcurrencyCounter struct {
	Counter int
	Mutex   sync.Mutex
}

func (c *ConcurrencyCounter) ConcurrencyIncrease() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Counter++
}

func (c *ConcurrencyCounter) Increase() {
	c.Counter++
}

func main() {
	cc := ConcurrencyCounter{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cc.ConcurrencyIncrease()
		}()
	}
	wg.Wait()
	fmt.Printf("Counter with mutex = %d\n", cc.Counter)

	cc.Counter = 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cc.Increase()
		}()
	}
	wg.Wait()
	fmt.Printf("Counter without mutex = %d\n", cc.Counter)

}
