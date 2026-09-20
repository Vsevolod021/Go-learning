package main

import "fmt"

func main() {
	// speciality := "Backend"
	name, lang := "Всеволод", "Go"
	age := 25
	isMan := true

	fmt.Println("Привет, Go-Компилятор!")
	fmt.Printf("Привет, %s\nДобро пожаловать в %s-разработку\n", name, lang)
	fmt.Printf("Я %v, мне %v лет\n", name, age)
	fmt.Printf("Я мужчина - %v\n", isMan)
}
