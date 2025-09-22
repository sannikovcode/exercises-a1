package main

import "fmt"

func main() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	if number >= 1 && number <= 10 {
		fmt.Printf("Ваше число: %d\n", number)
	} else {
		fmt.Printf("Ваше число: %d - не подходит к параметрам!\n", number)
		return
	}

	fmt.Printf("Таблица умнажения для числа:  %d от 1 до 10: \n", number)

	for i := 1; i <= 10; i++ {
		multiply := number * i
		fmt.Printf("%d X %d = %d\n", number, i, multiply)

	}

	fmt.Printf("Закончили расчет!")

}
