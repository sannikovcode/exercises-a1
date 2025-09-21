package main

import (
	"errors"
	"fmt"
	"strings"
)

func validateEmail(email string) error {
	atCount := strings.Count(email, "@")
	if atCount == 0 {
		return errors.New("Отсутсвует символ @")
	}
	if atCount > 1 {
		return errors.New("Слишком много символов @")
	}

	parts := strings.Split(email, "@")
	localPart := parts[0]
	domainPart := parts[1]

	if len(localPart) < 3 {
		return errors.New("Локадбная часть слишком короткая (минимум 3 символа)")
	}

	if len(domainPart) < 5 {
		return errors.New("Доменная часть слишком короткая (минимум 5 символов)")
	}

	if strings.HasPrefix(email, ".") {
		return errors.New("email не может начинаться с точки")
	}
	if strings.HasSuffix(email, ".") {
		return errors.New("email не может заканчиваться точкой")
	}
	return nil

}

func main() {
	var email string

	fmt.Println("Введите вашу почту: ")
	fmt.Scan(&email)

	if err := validateEmail(email); err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		fmt.Println("Попробуйте еще раз!")
	} else {
		fmt.Println("Email корректен!")
	}

}
