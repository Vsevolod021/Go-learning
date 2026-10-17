package main

import (
	"fmt"
	"runtime"
	"sync"
)

func say(wg *sync.WaitGroup, s string) {
	defer wg.Done()

	for range 3 {
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go say(&wg, "A")

	// fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	for range 100000 {
		wg.Go(func() {
			// time.Sleep(5 * time.Second)
			// fmt.Printf("Горутина %v\n", i)
			// fmt.Println(runtime.NumCPU())
			fmt.Println(runtime.NumGoroutine())
		})

	}

	fmt.Println(runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("все закончили")
	fmt.Println(runtime.NumGoroutine())
}
