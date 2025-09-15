package main

import (
	"fmt"
	"strings"
)

func main() {
	var email string

	fmt.Println("Введите ваше мыло: ")
	fmt.Scan(&email)

	hasOneAt := strings.Count(email, "@") == 1
	parts := strings.Split(email, "@")
	partLocalBefor := len(parts[0]) >= 3
	partAfter := len(parts[1]) >= 5

	hasStartDot := !strings.HasPrefix(email, ".")
	hasEndDOt := !strings.HasSuffix(email, ".")

	if hasOneAt && partLocalBefor && partAfter && hasStartDot && hasEndDOt {
		fmt.Println("Корректный")
	} else {
		fmt.Println("Некоретный")
	}

}
