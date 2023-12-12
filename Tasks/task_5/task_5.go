// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.
package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var timer int
	if _, err := fmt.Scan(&timer); err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	dataStream := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go DataReader(dataStream, done, &wg)
	go DataWriter(timer, dataStream, done, &wg)
	wg.Wait()
	close(dataStream)
}

// DataWriter В функции добавлен таймер, по истечении которого закрывает канал done и прекращается запись в dataStream
func DataWriter(timeout int, dataStream chan int, done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	timer := time.NewTimer(time.Duration(timeout) * time.Second)
	defer timer.Stop()

	for i := 0; ; i++ {
		select {
		case <-timer.C:
			println("\nStop writing data...")
			close(done)
			return
		default:
			dataStream <- i
		}
	}
}

func DataReader(dataStream chan int, done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		default:
			i := <-dataStream
			if i%10000000 == 0 {
				println("Reading data...")
			}
		}
	}
}
