package main

import (
	"fmt"
)

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a := 10
	b := 5

	fmt.Printf("До того как поменяли местами: %d и %d\n", a, b)

	swap(&a, &b)

	fmt.Printf("После того как поменяли местами: %d и %d\n", a, b)

}
