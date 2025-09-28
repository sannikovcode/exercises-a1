package main

import (
	"errors"
	"fmt"
	"strings"
)

func validateEmail(email string) error {
	if !strings.Contains(email, "@") {

		return errors.New("email должен содержать @")
	}
	parts := strings.Split(email, "@")

	if len(parts) != 2 {
		return errors.New("неверный формат email")
	}
	if !strings.Contains(parts[1], ".") {
		return errors.New("email должен содержать домен с точкой")
	}
	return nil

}

func main() {
	emails := []string{"user@example.com", "invalid-email", "user@com"}
	for _, email := range emails {
		err := validateEmail(email)
		if err != nil {
			fmt.Printf(" %s: %v\n", email, err)
		} else {
			fmt.Printf("%s: корректный email\n", email)
		}
	}
}
