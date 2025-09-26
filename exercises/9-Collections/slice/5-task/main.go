package main

import "fmt"

func main() {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}

	combined := append(first, second...)
	fmt.Println("Объединенный слайс:", combined)
}
