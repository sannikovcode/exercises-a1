package main

import "fmt"

func main() {
	var choice int
	var value float64

	fmt.Println("Выберите conversion: ")
	fmt.Println("1. Фарангейты в Цельсии")
	fmt.Println("2. Дюймы в сантиметры")
	fmt.Println("3. Фунты в килограммы")
	fmt.Scanln(&choice)

	fmt.Print("Введите значение для конвертации: ")
	fmt.Scanln(&value)

	switch choice {
	case 1:
		result := (value - 32) * 5 / 9
		fmt.Printf("%2f°F = %.2f°C\n", value, result)
	case 2:
		result := value * 2.54
		fmt.Printf("%.2f дюймов = %.2f см\n", value, result)
	case 3:
		result := value * 0.453592
		fmt.Printf("%.2f фунтов = %.2f кг\n", value, result)
	default:
		fmt.Println("Неизвестный выбор")
	}
}
