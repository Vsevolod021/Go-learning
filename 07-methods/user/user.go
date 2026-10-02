package user

import "github.com/Vsevolod021/Go-learning/07-methods/age"

type User struct {
	Name string
	Age  age.Age
}

func (u *User) Greet() string {
	return "Привет, " + u.Name
}

func (u *User) Birthday() {
	u.Age++
}
