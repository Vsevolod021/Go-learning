package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// ch := make(chan int, 10)

	// go generator(ch)

	// printAll(ch)

	// v, ok := <-ch

	// fmt.Printf("value - %#v\n", v)
	// fmt.Printf("is Ok - %v\n", ok)

	// fmt.Println()
	// fmt.Println("Несколько отправителей")

	var wg sync.WaitGroup

	results := make(chan int)

	numbers := []int{4, 12, 60}

	go func() {
		wg.Wait()
		close(results)
	}()

	for _, n := range numbers {
		wg.Go(func() { worker(results, n) })
	}

	for v := range results {
		fmt.Println(v)
	}

	// ch <- 42
}

func worker(ch chan<- int, n int) {
	ch <- n * n
}

func generator(ch chan<- int) {
	defer close(ch)

	for i := range 10 {
		ch <- i + 1
		fmt.Printf("отправил %v\n", i+1)
	}
}

func printAll(ch <-chan int) {
	for i := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(i)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func send(wg *sync.WaitGroup) {
	ch := make(chan int, 1)

	wg.Add(1)

	go func() {
		defer wg.Done()
		ch <- 42
	}()
}

func say(s string, ch chan struct{}) {
	for range 3 {
		fmt.Println(s)
	}
	fmt.Println("message before channel sent")
	ch <- struct{}{}

	fmt.Println("message after channel sent")
}
