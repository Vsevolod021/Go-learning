package models

type User struct {
	Name   string
	Age    int
	Email  string
	salary int // строчные поля неэкспортируемы
}
