package main

import "fmt"

func main() {

	number := 9

	if number%2 == 0 {
		fmt.Printf(" %d - четное число\n", number)

	} else {
		fmt.Printf("%d -нечетное чиисло\n", number)
	}
}
