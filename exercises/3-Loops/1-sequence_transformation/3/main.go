package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите число N: ")
	fmt.Scan(&n)

	fmt.Print("Числа: ")
	for i := 1; i <= n; i++ {
		if i > 1 {
			fmt.Print(", ")
		}
		fmt.Print(i)
	}
	fmt.Println()
}
