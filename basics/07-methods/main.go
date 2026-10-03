package main

import (
	"fmt"

	"github.com/Vsevolod021/Go-learning/basics/07-methods/age"
	"github.com/Vsevolod021/Go-learning/basics/07-methods/user"
)

func main() {
	u := user.User{Name: "Vsevolod", Age: 26}

	greeting := u.Greet()

	fmt.Println(greeting)

	fmt.Println("\nМутация")

	u.Birthday()

	fmt.Printf("%v\n", u.Age)

	(&u).Birthday()

	fmt.Printf("%v\n", u.Age)

	usersMap := map[string]*user.User{
		"Vsevolod": {Name: "Vsevolod", Age: 26},
		"Yulia":    {Name: "Yulia", Age: 27},
	}

	usersMap["Vsevolod"].Birthday()

	fmt.Printf("Возраст Всеволода - %v\n", usersMap["Vsevolod"].Age)

	usersMap["Maria"] = &user.User{Name: "Maria", Age: 2}

	fmt.Printf("Взрослая ли Юлия - %v\n", usersMap["Yulia"].Age.IsAdult())
	fmt.Printf("Взрослая ли Маша - %v\n", usersMap["Maria"].Age.IsAdult())

	// usersMap2 := map[string]user.User{
	// 	"Vsevolod": {Name: "Vsevolod", Age: 26},
	// 	"Yulia":    {Name: "Yulia", Age: 27},
	// }

	// usersMap2["Yulia"].Birthday() // cannot call pointer method Birthday on user.User

	const n = 5 // Нетипизированная константа

	fmt.Println(age.Age(20) + n)
}
