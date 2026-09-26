package main

import "fmt"

func birthday(age *int) {
	*age++
}

// Для работы с числовым значением
func main() {
	age := 25

	fmt.Printf("%v\n", age)
	fmt.Println("Incrementing...")

	p := &age

	birthday(p)

	fmt.Printf("%v\n", age)

	// q := &age // Взятие адреса от age
	// var q *int // Объявление nil-указателя
	q := new(int) // Объявление указателя с new. Значение - zero value для указанного типа данных

	// *q = 10

	if q == nil {
		fmt.Println("Не передано значение")
		return
	}
	fmt.Printf("Значение равно %v\n", *q)

}
