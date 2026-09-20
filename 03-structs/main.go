package main

import (
	"fmt"

	"github.com/Vsevolod021/Go-learning/03-structs/models"
)

type Job struct {
	position string
	salary   int
}

func birthday(user models.User) {
	user.Age = 30

	fmt.Println(user.Age) // 30
}

func main() {
	user1 := models.User{Name: "Vsevolod", Age: 25, Email: "seva@gmail.com"}
	// var user2 models.User

	birthday(user1)

	fmt.Println(user1.Age) // 25

	// fmt.Printf("%#v\n", user1)   // main.User{Name:"Vsevolod", Age:25, Email:"seva@gmail.com"}
	// fmt.Printf("%#v\n\n", user2) // main.User{Name:"", Age:0, Email:""}

	// fmt.Printf("%v\n", user1)    // {Name:Vsevolod Age:25 Email:seva@gmail.com}
	// fmt.Printf("%+v\n\n", user2) // {Name: Age:0 Email:}

	// jsonUser1, err := json.Marshal(user1)

	// if err != nil {
	// 	log.Fatalf("Ошибка кодирования jsonUser1")
	// }
	// fmt.Printf("%v\n", string(jsonUser1))

	// jsonUser2, err := json.Marshal(user2)

	// if err != nil {
	// 	log.Fatalf("Ошибка кодирования jsonUser2")

	// }
	// fmt.Printf("%s\n", jsonUser2)

	// fmt.Printf("%-v\n", user1)   // {Vsevolod 25 seva@gmail.com}
	// fmt.Printf("%-v\n\n", user2) // { 0 }

	// fmt.Printf("%0v\n", user1)   // {Vsevolod 25 seva@gmail.com}
	// fmt.Printf("%0v\n\n", user2) // { 0 }

	// fmt.Printf("% v\n", user1)   // {Vsevolod  25 seva@gmail.com}
	// fmt.Printf("% v\n\n", user2) // {  0 }

	// job1 := Job{position: "Backend Go", salary: 600000}

	// fmt.Printf("%#v\n", job1)
}
