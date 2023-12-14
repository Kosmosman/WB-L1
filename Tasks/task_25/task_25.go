package main

import (
	"fmt"
	"time"
)

// Создаем новый объект Ticker, который по истечение t секунд будет отправлять в канал текущее время.
// Поскольку это функция Sleep, мы выходим из цикла ожидания поступления сигнала после первого тика
func SleepTick(t float64) {
	c := time.NewTicker(time.Duration(t) * time.Second)
	defer c.Stop()
	for range c.C {
		fmt.Printf("Time is up after tick in %v!\n", time.Now())
		break
	}
}

// Запоминаем время, когда пройдёт t секунд, и в цикле проверяем, достигли ли мы этого момента
func SleepAfterAddTime(t float64) {
	end := time.Now().Add(time.Duration(t) * time.Second)
	for time.Now().Before(end) {
		continue
	}
	fmt.Printf("Time is up after add time in %v!\n", time.Now())
}

// Создаем канал, который по истечение t секунд будет посылать сигнал
func SleepTimeAfter(t float64) {
	<-time.After(time.Duration(t) * time.Second)
	fmt.Printf("Time is up after timeAfter in %v!\n", time.Now())
}

func main() {
	fmt.Printf("Start sleep at %v\n", time.Now())
	SleepTick(2)
	fmt.Printf("Start sleep at %v\n", time.Now())
	SleepAfterAddTime(2)
	fmt.Printf("Start sleep at %v\n", time.Now())
	SleepTimeAfter(2)
}
