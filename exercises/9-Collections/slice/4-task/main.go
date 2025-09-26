package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("1. Исходный слайс:", fruits)

	fruits = append(fruits, "date")
	fmt.Println("2. После добавления `date` :", fruits)

	fmt.Println("3. fruits[:1] =", fruits[:1])
	fmt.Println("4. fruits[2:] =", fruits[2:])

	fruits = append(fruits[:1], fruits[2:]...)
	fmt.Println("5. После удаления элемента [1]:", fruits)

}
