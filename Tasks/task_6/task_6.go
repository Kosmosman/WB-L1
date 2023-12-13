// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	stopChanel := make(chan struct{})
	go GoroutineStopWithChan(stopChanel, &wg)
	time.Sleep(2 * time.Second)
	close(stopChanel)

	var stopFlag atomic.Bool
	go GoroutineStopWithFlag(&stopFlag, &wg)
	time.Sleep(2 * time.Second)
	stopFlag.Swap(true)

	ctx, cancel := context.WithCancel(context.Background())
	go GoroutineStopWithContext(ctx, &wg)
	time.Sleep(2 * time.Second)
	cancel()

	wg.Wait()
}

// Остановка горутины с помощью закрытия канала
func GoroutineStopWithChan(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ch:
			println("End goroutine with signal chanel")
			return
		default:
			continue
		}
	}
}

// Остановка горутины с помощью завершения контекста
func GoroutineStopWithContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			println("End goroutine with context done")
			return
		default:
			continue
		}
	}
}

// Остановка горутины с передачей внешнего флага
func GoroutineStopWithFlag(flag *atomic.Bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for !flag.Load() {
	}
	println("End goroutine with foreign flag")
}
