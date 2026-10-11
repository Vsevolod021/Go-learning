package main

import (
	"fmt"
	"os"
)

func order() {
	fmt.Println("start")
	defer fmt.Println("deferred")
	defer fmt.Println("1..")
	defer fmt.Println("2..")
	defer fmt.Println("3...")
	fmt.Println("end")
}

func deferArgsEvaluation() {
	x := 1
	// defer fmt.Println("defer", x) // запоминает аргумент на момент

	defer func() { fmt.Println("defer func no args", x) }() // привязывается

	if true {
		x = 2
		return
	}

	defer func(x int) { fmt.Println("defer func with arg", x) }(x)
	fmt.Println("now", x)
}

func writeFile(name, text string) error {
	f, err := os.Create(name)

	if err != nil {
		return fmt.Errorf("create file %v: %w", name, err)
	}
	defer f.Close()

	_, err = f.WriteString(text)

	return err
}

func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("divide: %v / %v: %v", a, b, r)
		}
	}()

	return a / b, nil
}

func main() {
	order()
	deferArgsEvaluation()

	err := writeFile("defer-study1.txt", "Hello! It's Vsevolod. Go developer")

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Println()
	fmt.Println("panic и recover")

	values := [][2]int{{10, 2}, {10, 0}}

	for _, v := range values {
		result, err := safeDivide(v[0], v[1])

		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		fmt.Println(result)
	}

}
