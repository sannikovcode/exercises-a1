package main

import (
	"errors"
	"fmt"
)

func validatePassword(password string) (bool, error) {
	if len(password) == 0 {
		return false, errors.New("пароль не может быть пустым")
	}
	if len(password) < 8 {
		return false, errors.New("пароль должен быть не менее 8 символов")
	}
	return true, nil
}

func main() {
	passwords := []string{"", "123", "strongpassword"}

	fmt.Println("Проверка паролей:")
	fmt.Println("---------------------")

	for _, pwd := range passwords {
		isValid, err := validatePassword(pwd)
		if err != nil {
			fmt.Printf("Пароль '%s': %v\n", pwd, err)
		} else {
			fmt.Printf("Пароль '%s': пройден(правильно: %t)\n", pwd, isValid)
		}
	}

}
