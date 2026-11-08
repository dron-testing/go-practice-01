package main

import "fmt"

func main() {
	var name string = "Саша"
	var group string = "ЭФБО-03-26"
	const cours string = "Бэкенд-разработка на языке Go"
	completed := 2
	var intersted bool = true

	fmt.Printf("Участник: %s\n", name)
	fmt.Printf("Группа: %s\n", group)
	fmt.Printf("Курс: %s\n", cours)
	fmt.Printf("Завершено работ: %d\n", completed)
	fmt.Printf("Интерес к бэкенду: %t час\n", intersted)
}
