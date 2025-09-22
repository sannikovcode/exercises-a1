package main

import "fmt"

func main() {

	x := 10

	p := &x

	fmt.Printf("значение x: %d\n", x)
	fmt.Printf("адерес x: %d\n", &x)
	fmt.Printf("значение указателя p: %d\n", p)
	fmt.Printf("значение по адресу в указателе: %d\n", *p)

	*p = 20

	fmt.Printf("значение x после изменения: %d\n", *p)

}
