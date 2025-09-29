package main

import (
	"errors"
	"fmt"
)

func validateAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}
	if age > 150 {
		return errors.New("age seems unrealistic")
	}
	return nil
}
func main() {
	if err := validateAge(-5); err != nil {
		fmt.Println("Validation error:", err)
	}
}
