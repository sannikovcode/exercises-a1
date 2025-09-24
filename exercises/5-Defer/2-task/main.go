package main

import "fmt"

func main() {
	x := 10

	defer fmt.Println("Defer:", x)
	x = 20

	fmt.Println("Main:", x)
}
