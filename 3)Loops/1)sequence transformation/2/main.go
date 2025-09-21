package main

import "fmt"

func main() {

	var n int

	fmt.Print("Введите число N: ")
	_, err := fmt.Scan(&n)

	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	if n < 1 {
		fmt.Println("N должно быть больше или равно 1")
		return
	}

	fmt.Printf("Числа от 1 до %d: \n", n)
	for i := 1; i <= n; i++ {

		fmt.Println(i)
	}
}
