package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		log.Fatal(err)
	}
	dataStream := make(chan int) // Канал на передачу и чтение иформации
	done := make(chan struct{})  // Канал на закрытие горутин

	go Listener(done)
	for i := 0; i < n; i++ {
		go DataPrinter(dataStream, done)
	}
	DataWriter(dataStream, done)
	close(dataStream)
}

// Listener Подпрограмма, ожидающая ^C сигнала. После получения закрывает канал, сигнализирующий об остановке остальных горутин
func Listener(done chan struct{}) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT)
	<-sig
	fmt.Println("\nReceived SIGINT, exiting...")
	close(done)
}

// DataPrinter Выводит на перать данные, полученные из канала dataStream
func DataPrinter(dataStream chan int, done chan struct{}) {
	for {
		select {
		case <-done:
			return
		default:
			println(<-dataStream)
		}
	}
}

// DataWriter Передает данные в канал dataStream
func DataWriter(dataStream chan int, done chan struct{}) {
	for i := 0; ; i++ {
		select {
		case <-done:
			return
		default:
			dataStream <- i
		}
	}
}
