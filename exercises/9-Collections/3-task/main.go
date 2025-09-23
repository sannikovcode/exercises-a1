package main

import "fmt"

func main() {
	temperatures := [7]int{15, 18, 12, 20, 25, 22, 16}
	daysOfWeek := [7]string{"Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота", "Воскресенье"}

	sum := 0
	max := temperatures[0]
	min := temperatures[0]
	daysAboveAverage := 0

	fmt.Println("Температура по дням неделям:")

	for i, value := range temperatures {

		fmt.Printf("%s : %d\n", daysOfWeek[i], value)

		sum += value

		if value > max {
			max = value
		}
		if value < min {
			min = value
		}

	}
	result := float64(sum) / float64(7)

	for _, value := range temperatures {
		if float64(value) > result {
			daysAboveAverage++
		}

	}
	fmt.Println("-------------------")

	fmt.Printf("Средняя температура: %.2fC\n", result)
	fmt.Printf("Максимальная температура: %d\n", max)
	fmt.Printf("Минимальная температура: %d\n", min)
	fmt.Printf("Дней с температурой выше средней: %d\n", daysAboveAverage)

}
