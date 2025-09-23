package main

import "fmt"

func main() {

	numbers := [6]int{5, 13, -4, 7, 0, 11}

	sum := 0
	max := numbers[0]

	for i, value := range numbers {

		fmt.Printf("Значения индекса: [%d] = %d\n ", i, value)
		sum += value

		if value > max {
			max = value
		}

	}
	fmt.Println("Сумма всех элементов:", sum)

	fmt.Println("Максимальный элемент:", max)

}
