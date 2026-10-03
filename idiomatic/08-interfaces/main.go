package main

import "fmt"

type Greeter interface {
	Greet() string
}

// type Stringer interface {
// 	String() string
// }

type Bot struct {
	Name string
}
type User struct {
	Name string
}
type Cat struct {
	Name string
}
type Dog struct {
	Name string
}

func (b Bot) Greet() string {
	return "BEEP. Unit " + b.Name + " online"
}

func (u *User) Greet() string {
	return "Приветствую, " + u.Name
}

func (c Cat) Greet() string {
	return "Приветствую, " + c.Name
}
func (d Dog) String() string {
	return fmt.Sprintf("Кличка собаки - %v", d.Name)
}

func sayHello(g Greeter) {
	// t := g.(type) // invalid syntax tree: use of .(type) outside type switch

	switch v := g.(type) {
	case *User:
		fmt.Printf("%T\n", v) // *main.User
	case Bot:
		fmt.Printf("%T\n", v) // main.Bot
	default:
		fmt.Printf("%T\n", v)
		// fmt.Println(v.Name) // v.Name undefined (type Greeter has no field or method Name)
	}
	fmt.Println(g.Greet())
}

func describe(x any) {
	fmt.Printf("\n%T\n", x)

	switch v := x.(type) {
	case int:
		fmt.Printf("число %d\n", v*2)
	case Greeter:
		fmt.Println(v.Greet())
	case string:
		fmt.Printf("строка %v\n", len(v))
	default:
		fmt.Printf("Неизвестный тип - %T\n", v)
	}
}

func main() {
	u := User{Name: "Vsevolod"}
	b := Bot{Name: "ClaudeAI"}
	c := Cat{Name: "Barsik"}

	sayHello(&u)
	sayHello(b)
	sayHello(&c)

	fmt.Println()
	fmt.Println("nil-указатель")

	var p *User = nil
	var g Greeter = p

	fmt.Println(p == nil)
	fmt.Println(g == nil)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", g)

	fmt.Println()
	fmt.Println("Сужение типа")

	describe(42)     // сложил
	describe("text") // вывел длину
	describe(&u)     // поприветствовал
	describe(u)      // не определился как Greeter, вероятнее всего потому, что Greeter есть у указателя на User, а не у самого User
	describe(b)      // поприветствовал
	describe(nil)    // не определился как не default

	fmt.Println()
	fmt.Println("Интерфейс из стандартной библиотеки")

	d := Dog{Name: "Sparky"}
	fmt.Println(&d)
	fmt.Println(d)
	// fmt.Printf("%v\n", d)
}
