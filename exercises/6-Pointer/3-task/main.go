package main

import "fmt"

func main() {
	number := 100
	pointer := &number
	pointer2 := &number

	fmt.Printf("Значение number: %d\n", number)
	fmt.Printf("Значение через pointer: %d\n", *pointer)
	fmt.Printf("Значение через pointer2: %d\n", *pointer2)

	*pointer2 = 200

	fmt.Printf("Значение number: %d\n", number)
	fmt.Printf("Значение через pointer: %d\n", *pointer)
	fmt.Printf("Значение через pointer2: %d\n", *pointer2)

}
