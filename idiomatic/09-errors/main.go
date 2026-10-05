package main

import (
	"errors"
	"fmt"
	"strconv"
)

type AgeError struct {
	Value int
	Max   int
}

const maxAge int = 150

func (e *AgeError) Error() string {
	return fmt.Sprintf("age %v exceeds max %v", e.Value, e.Max)
}

var ErrNegativeAge = errors.New("возраст не может быть меньше нуля")

func parseAge(s string) (int, error) {
	x, err := strconv.Atoi(s)

	if err != nil {
		return 0, fmt.Errorf("parse age %q: %w", s, err)
	}

	if x < 0 {
		return 0, fmt.Errorf("parse age %q: %w", s, ErrNegativeAge)
	}

	if x > maxAge {
		return 0, fmt.Errorf("parse age %q: %w", s, &AgeError{Value: x, Max: maxAge})
	}

	return x, nil
}

func check(age int) error {
	if age > maxAge {
		return &AgeError{Value: age, Max: maxAge}

	}
	return nil
}

func main() {
	values := []string{"25", "abc", "-5", "151"}

	for _, v := range values {
		a, err := parseAge(v)

		if err != nil {
			fmt.Println()

			if errors.Is(err, ErrNegativeAge) {
				fmt.Printf("%v. Введите положительное число\n", err)
				continue
			}

			var numErr *strconv.NumError

			if errors.As(err, &numErr) {
				fmt.Printf("%v\n", numErr.Err)
				fmt.Printf("%v\n", err)

				// log.Fatalf("%v\n", err) // Завершает выполнение программы
				// fmt.Printf("После") // Уже не выполнится

				continue
			}

			var ageErr *AgeError

			if errors.As(err, &ageErr) {
				fmt.Printf("Указан возраст %v, это на %v больше максимума %v\n", ageErr.Value, ageErr.Value-ageErr.Max, ageErr.Max)
				fmt.Printf("%v\n", err)
				continue
			}

			fmt.Printf("%v\n", err)
			continue
		}
		fmt.Printf("Возраст равен %v\n", a)
	}

	fmt.Println()
	fmt.Println("nil-ловушка")
	err := check(20)

	fmt.Printf("%v\n", err == nil)
}
