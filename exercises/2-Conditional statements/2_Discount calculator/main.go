package main

import "fmt"

func main() {

	sum := 44.0
	if sum > 1000 {
		discount := sum * 0.10
		sum = sum - discount
	}
	fmt.Printf("Итоговая сумма: %.0f руб.\n", sum)

}
