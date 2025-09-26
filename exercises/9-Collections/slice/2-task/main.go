package main

import "fmt"

func main() {
	numbers := []int{34, 7, 23, 32, 5, 62}
	max := numbers[0]
	for _, num := range numbers[1:] {
		if num > max {
			max = num
		}
	}
	fmt.Println("Максимальный элемент:", max)

}
