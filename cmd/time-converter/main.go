package main

import "fmt"

func main() {
	var minutes int
	const minInHour int = 60
	fmt.Print("Введите количество минут: ")
	_, err := fmt.Scan(&minutes)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if minutes < 0 {
		fmt.Println("Ошибка: число не должно быть отрицательным.")
		return
	}
	fmt.Printf("Полных часов: %d\n", minutes/minInHour)
	fmt.Printf("Оставшихся минут: %d\n", minutes%minInHour)
}
