package main

import (
	"fmt"

	"github.com/Vsevolod021/Go-learning/03-structs/models"
)

// Для работы со структурой
func birthday(u *models.User) {
	// (*u).Age++ // явное разыменовывание
	u.Age++ // синтаксический сахар
}

func main() {
	user := models.User{Name: "Vsevolod", Age: 25, Email: "Seva@gmail.com"}

	fmt.Printf("%v\n", user.Age)
	fmt.Println("Incrementing...")

	birthday(&user)

	fmt.Printf("%v\n", user.Age)
}
