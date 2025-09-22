package main

import "fmt"

func main() {
	var number int
	fmt.Println("Введите число: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Printf("Ваше число: %d - четное\n", number)

	} else {
		fmt.Printf("Ваше число: %d - не четное\n", number)
	}

}