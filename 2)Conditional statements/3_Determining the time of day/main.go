package main

import "fmt"

func main() {

	hour := 6

	if hour >= 6 && hour <= 11 {
		fmt.Println("Сейчас утро")

	} else if hour >= 12 && hour <= 17 {
		fmt.Println("Сейчас день")

	} else if hour >= 18 && hour <= 22 {
		fmt.Println("Сейчас вечер")
	} else {
		fmt.Println("Сейчас ночь")
	}
}
