package main

import "fmt"

func main() {

	var start, end, step int

	fmt.Print("Введите начальное число:", start)
	fmt.Scan(&start)

	fmt.Print("Введите конечное число:", end)
	fmt.Scan(&end)

	fmt.Print("Введите значения шага:", step)
	fmt.Scan(&step)

	if step <= 0 {
		fmt.Println("Шаг должен быть положительным числом!")
		return
	}

	if start <= end {
		isFirst := true
		for i := start; i <= end; i += step {
			if !isFirst {
				fmt.Print(", ")
			}
			fmt.Print(i)
			isFirst = false

		}
	} else {
		isFirst := true
		for i := start; i >= end; i -= step {
			if !isFirst {
				fmt.Print(", ")
			}
			fmt.Print(i)
			isFirst = false
		}
	}
	fmt.Println()

}
