package main

import "fmt"

func main() {

	name := "GO"
	ptr := &name

	fmt.Printf("значения name: %s\n", name)

	fmt.Printf("адрес  name: %p\n", &name)
	fmt.Printf("значение указателя ptr: %p\n", ptr)
	fmt.Printf("значение по адресу через ptr: %s\n", *ptr)

	*ptr = "Golang"

	fmt.Printf("значения name после изменения: %s\n", name)

}
