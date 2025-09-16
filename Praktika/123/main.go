package main

import "fmt"

func main() {
	var x int

	fmt.Println("Введите пожалуйста число: ")
	_, err := fmt.Scan(&x)

	if err != nil {
		fmt.Println("Ошибка! Введите число: ")
		return
	}

	if x <= 0 {
		fmt.Println("Ошибка! Введите пожалйста положительное число!")
		return
	}

	fmt.Println("Считаем от 1 до: ", x, ":")

	for i := 1; i <= x; i++ {
		fmt.Println(i)
	}

	fmt.Println("Кончили!")

}
