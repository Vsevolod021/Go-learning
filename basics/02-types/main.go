package main

import (
	"fmt"
)

var someNumber int
var someFloat float64
var someString string
var someBool bool

func main() {
	// a := 10
	// b := 2.5

	// str1 := "Строка 1"
	// str2 := "Строка 2"

	// fmt.Printf("%v, %v, %v, %v\n", someNumber, someFloat, someString, someBool)
	// fmt.Printf("%T, %T, %T, %T\n", someNumber, someFloat, someString, someBool)
	// fmt.Printf("%v\n", float64(a)+b)
	// fmt.Printf("%v\n", str1+str2) // Конкатенация как в JS, хехе

	// c := -2.5

	// fmt.Println(int(c))
	// fmt.Println(len("Привет, Го"))

	greeting := "Привет, Go"
	count := 0

	for range greeting {
		count++
	}

	for _, r := range greeting {
		fmt.Printf("%c\n", r)
		count++
	}

	// fmt.Println(count)
	// fmt.Println(utf8.RuneCountInString(greeting))
}
