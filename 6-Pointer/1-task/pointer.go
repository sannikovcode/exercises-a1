package main

import "fmt"

func main() {

	x := 10

	p := &x

	fmt.Printf("Значение x: %d\n", x)
	fmt.Printf("адрес x: %d\n", &x)
	fmt.Printf(" значения указателя: %d\n", p)
	fmt.Printf("значения по адресу: %d\n", *p)

	*p = 20

	fmt.Printf("значение x после изменения: %d\n", *p)

}
