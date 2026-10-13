package main

import (
	"cmp"
	"fmt"
)

type Number interface {
	~int | ~float64
}

type Age int

func Sum[T Number](nums []T) (result T) {
	for _, num := range nums {
		result += num
	}
	return result
}

func Contains[T comparable](s []T, v T) bool {
	for _, value := range s {
		if value == v {
			return true
		}
	}
	return false
}

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return last, true
}

func (s *Stack[T]) Map[U any](f func(T) U) []U {
	var u = make([]U, 0, len(s.items))

	for _, v := range s.items {
		u = append(u, f(v))
	}
	return u
}

func PrintAll[T fmt.Stringer](items []T) {
	for _, v := range items {
		fmt.Printf("%v", v)
	}
}

func printOne(item fmt.Stringer) {
	fmt.Print(item)
}

func (d Dog) String() string {
	return fmt.Sprintf("%v barking and greets you\n", d.Name)
}

func (c Cat) String() string {
	return fmt.Sprintf("%v meowing and greets you\n", c.Name)
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

type Pet interface {
	String() string
}

func main() {
	intNums := []int{1, 50, 77, 61}
	float64Nums := []float64{1, 55.5, 34.2, 11.1}

	fmt.Println(Sum(intNums))
	fmt.Println(Sum(float64Nums))

	ages := []Age{26, 25}

	fmt.Println(Sum(ages))

	fmt.Println(Contains(intNums, 50))
	fmt.Println(Contains(intNums, 51))

	names := []string{"Seva", "Yulia", "Masha"}

	fmt.Println(Contains(names, "Masha"))
	fmt.Println(Contains(names, "Volodya"))

	fmt.Println(Max(10, 2))
	fmt.Println(Max("Vsevolod", "Yulia"))

	// s := Stack[int]{items: []int{26, 25, 1, 2}}
	// s.Push(4)

	s := Stack[string]{items: []string{"Vsevolod", "Yulia", "Masha"}}
	s.Push("Volodya")

	fmt.Printf("%v\n", s.items)

	// fmt.Println()
	// fmt.Println("Pop values")

	// for range len(s.items) + 1 {
	// 	popItem, popSuccess := s.Pop()

	// 	fmt.Printf("pop item: %#v, pop success: %v, items left: %v\n", popItem, popSuccess, s.items)
	// }

	fmt.Println()
	fmt.Println("Map")

	lens := s.Map(func(v string) int {
		return len(v)
	})

	fmt.Println(lens)

	fmt.Println()
	fmt.Println("last question")

	cat1 := Cat{Name: "Barsik"}
	cat2 := Cat{Name: "Sima"}
	cat3 := Cat{Name: "Begemot"}

	dog3 := Dog{Name: "Sparky"}
	dog4 := Dog{Name: "Spike"}

	pets := []Pet{cat1, cat2, cat3, dog3, dog4}

	// PrintAll(cats)

	for _, p := range pets {
		printOne(p)
	}
}
