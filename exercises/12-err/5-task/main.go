package main

import (
	"errors"
	"fmt"
)

func calculateAge(birthYear, currentYear int) (int, error) {
	if birthYear < 0 {
		return 0, errors.New("год рождения не может быть отрицательным")
	}
	if birthYear > currentYear {
		return 0, errors.New("год рождения не может быть больше текущего года")
	}
	return currentYear - birthYear, nil
}

func main() {
	age, err := calculateAge(1990, 2024)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Возраст:", age)
	//--------------------------------------------------------------------

	age, err = calculateAge(2030, 2024)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Возраст:", age)
	}

}
