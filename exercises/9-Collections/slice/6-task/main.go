package main

import "fmt"

func main() {

	numbers := make([]int, 3, 5)
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30

	numbers = append(numbers, 40, 50)
	fmt.Println("Слайс:", numbers)
	fmt.Println("Длина:", len(numbers))
	fmt.Println("Вместимость:", cap(numbers))
}
