package main

import (
	"fmt"
)

// Address
type Address struct {
	City   string
	Street string
	Name   string
}

func (a Address) Locate() string {
	return fmt.Sprintf("Я нахожусь по адресу: город %v, улица %v, дом %v", a.City, a.Street, a.Name)
}

// Person
type Person struct {
	Name string
	Age  int
}

func (p Person) Title() string {
	return "человек"
}

func (p Person) Greet() string {
	return fmt.Sprintf("Привет, меня зовут %v и я %v", p.Name, p.Title())
}

func introduce(p Person) {
	fmt.Printf("Приветствую, меня зовут %v\n", p.Name)
}

// Employee
type Employee struct {
	Person // встраивание: указан только тип, без имени поля
	Address
	Position string
}

func (e Employee) Title() string {
	return "сотрудник"
}

func (e Employee) Greet() string {
	return fmt.Sprintf("Привет, меня зовут %v и я %v", e.Person.Name, e.Title())
}

// Интерфейсы
type Greeter interface {
	Greet() string
}

type Locator interface {
	Locate() string
}

type Profile interface {
	Greeter
	Locator
}

func getProfile(p Profile) string {
	return fmt.Sprintf("%v\n%v\n", p.Greet(), p.Locate())
}

func sayHello(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	// 1
	e := Employee{
		Position: "Backend",
		Person:   Person{Name: "Vsevolod", Age: 26},
		Address:  Address{City: "Moscow", Street: "Kutuzovsky", Name: "94/1"},
	}

	// 2
	// fmt.Println(e.Name)
	fmt.Println(e.Person.Name)

	// 3
	fmt.Println(e.Greet())
	fmt.Println(e.Person.Greet())

	fmt.Println(e.Locate())
	fmt.Println(e.Address.Locate())

	// 4
	fmt.Printf("%+v\n", e)

	introduce(e.Person)

	fmt.Println()
	fmt.Println("Интерфейсы")

	sayHello(e)
	sayHello(e.Person)

	fmt.Println()
	fmt.Println("Конфликт имен")

	// fmt.Println(e.Name)
	fmt.Printf("%#v\n", e.Address.Name)
	fmt.Println(e.Street)

	fmt.Println()
	fmt.Println("Встраиваемые интерфейсы")
	fmt.Println(getProfile(e))
}
