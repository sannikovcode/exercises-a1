package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "cherry"}
	fruits = append(fruits, "date")

	fruits = append(fruits[:1], fruits[2:]...)
	fmt.Println("Итоговый слайс:", fruits)
}
