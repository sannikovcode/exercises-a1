package main

import "fmt"

func main() {
	grades := [7]int{5, 7, 8, 6, 9, 10, 4}

	daysOfWeek := [7]string{"Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота", "Воскресенье"}

	sum := 0
	maxGrade := grades[0]

	fmt.Println("Журнали оценок студента: ")

	for i, grade := range grades {
		fmt.Printf("%s: %d\n", daysOfWeek[i], grade)

		sum += grade

		if grade > maxGrade {
			maxGrade = grade

		}

	}
	average := float64(sum) / float64(len(grades))

	fmt.Println("---------------------------")

	fmt.Printf("Средний балл студента: %.2f\n ", average)
	fmt.Printf("Лучшая оценка: %d\n", maxGrade)
}
