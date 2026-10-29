package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// ticker := time.NewTicker(100 * time.Millisecond)
	// defer ticker.Stop()

	// done := make(chan struct{})

	// go func() {
	// 	time.Sleep(550 * time.Millisecond)
	// 	close(done)
	// }()

	// tickerHandler(ticker, done)

	s1 := []int{1, 2, 3}
	s2 := []int{10, 20, 30, 40, 50}

	ch1 := make(chan int, len(s1))
	ch2 := make(chan int, len(s2))

	go generator(ch1, s1)
	go generator(ch2, s2)

	for ch1 != nil || ch2 != nil {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				break
			}
			fmt.Printf("value - %#v, ok - %v\n", v, ok)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			fmt.Printf("value - %#v, ok - %v\n", v, ok)
		}
	}

	fmt.Println(runtime.NumGoroutine())
}

func tickerHandler(ticker *time.Ticker, done <-chan struct{}) {
	for {
		select {
		case <-ticker.C:
			fmt.Println("Тик")
		case <-done:
			fmt.Println("Стоп")
			return
		}
	}
}

func generator[T any](ch chan<- T, s []T) {
	defer close(ch)

	for i := range s {
		ch <- s[i]
	}
}
