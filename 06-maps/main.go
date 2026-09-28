package main

import (
	"fmt"
	"maps"
	"slices"
)

func mutateMap(m map[string]int) {
	m["Begemot"] = 12
}

func main() {
	ages := map[string]int{
		"Vsevolod": 25,
		"Anna":     30,
	}

	fmt.Println("\nБазовые операции с map...")

	fmt.Printf("%T\n", ages) // map[string]int
	fmt.Printf("%v\n", ages) // map[Anna:30 Vsevolod:25]
	fmt.Println(len(ages))   // 2

	ages["Vsevolod"] = 26 // Изменение существующего
	ages["Yulia"] = 27    // Добавление нового
	delete(ages, "Anna")  // Удалениe старого

	fmt.Printf("%v\n", ages)
	fmt.Printf("%v\n", ages["Sonya"]) // Не падает - Zero Value

	fmt.Println("\nСоздание через make...")

	professions := make(map[string]string)

	fmt.Printf("%v\n", professions) // map[]

	fmt.Println("\nTwo values assignment...")

	v, ok := professions["Vsevolod"]

	fmt.Printf("%#v\n", v) // Zero Value
	fmt.Printf("%v\n", ok) // false. Не существует такого ключа

	professions["Vsevolod"] = "Backend"

	if v, ok := professions["Vsevolod"]; ok {
		fmt.Printf("%v\n", v)  // Backend
		fmt.Printf("%v\n", ok) // true. Ключ есть
	}

	var m map[string]int

	fmt.Println("\nNil-мапа")

	if m == nil {
		fmt.Println("мапа m - nil")
		fmt.Printf("%v\n", len(m))        // 0
		fmt.Printf("%v\n", m["Vsevolod"]) // Zero Value

		// m["Vsevolod"] = 26 // panic: assignment to entry in nil map

		m = make(map[string]int)
	}

	m["Vsevolod"] = 26
	fmt.Printf("%v\n", m["Vsevolod"]) // Не падает, потому что в блоке if была создана мапа

	fmt.Println("\nПорядок обхода")

	// ages["Barbara"] = 22 // Добавил еще одно поле чтобы проверить сортировку
	// Хочу быть с Юлей вдвоем, поэтому закомментил :)))

	for range 6 {
		for k, v := range ages {
			fmt.Printf("%v - %v, ", k, v) // Я намеренно не добавил \n
		}
		fmt.Println()
		// Мы получили, что то Юлия первая, то Всеволод
	}

	// отсортированные ключи
	keys := make([]string, 0, len(ages))

	for k := range ages {
		keys = append(keys, k)
	}
	// sort.Strings(keys) // Через пакет sort
	slices.Sort(keys) // Через пакет slices

	fmt.Println("\nВывод отсортированные ключи")

	for range 6 {
		for _, v := range keys {
			fmt.Printf("%v - %v, ", v, ages[v]) // Я намеренно не добавил \n
		}
		fmt.Println()
	}

	// Сравнение мап
	fmt.Println("\nСравнение мап")

	ages2 := maps.Clone(ages)

	fmt.Printf("%v\n", ages2)
	// fmt.Printf("%v\n", ages == ages2) // invalid operation: ages == ages2 (map can only be compared to nil)

	fmt.Printf("%v\n", maps.Equal(ages, ages2))

	ages2["Rex"] = 4

	fmt.Printf("%v\n", maps.Equal(ages, ages2))

	fmt.Println("\nМапа в функции")
	mutateMap(ages2)

	fmt.Printf("%v\n", ages2) // Гипотеза: мутировал. Факт: Мутировал
}
