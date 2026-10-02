package main

import (
	"fmt"
	"strconv"
)

func add(s []int) []int {
	return append(s, 99)
}

func main() {
	// Часть 1. Массивы
	const size = 3

	var arr [size]string // size обязательно константа

	for i := range arr {
		arr[i] = fmt.Sprintf("string%d", i+1) // 1 способ: Sprintf
		// arr[i] = "string" + strconv.Itoa(i+1) // 2 способ: strconv
		// arr[i] = "string" + string(i+1) // go vet ругается
	}

	fmt.Printf("%T\n", arr)
	fmt.Printf("%v\n", arr)

	// Часть 2. Слайсы
	s := []string{"a", "b", "c", "d"}

	fmt.Printf("%T\n", s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	// append(s, "d") // Ошибка

	for i := range 19 {
		fmt.Printf("Итерация %d...\n", i+1)
		s = append(s, strconv.Itoa(i))

		fmt.Println(len(s))
		fmt.Println(cap(s))
	}

	fmt.Println("\nКопирование слайса")

	e := []int{1, 2, 3, 4, 5}

	f := make([]int, len(e))
	copy(f, e)

	fmt.Printf("%v\n", f)

	e = add(e)

	fmt.Printf("%v\n", e)

	var n []int

	if n == nil {
		fmt.Println("n - это nil")
	}
}
