package main

import "fmt"

func main() {
	var number int
	fmt.Print("Введите число положительное и целое число: ")
	fmt.Scan(&number)


     
	sum := 0
    for i := 1; i <= number; i++ {
		sum += i
}
	fmt.Printf("Сумма чисел от 1 до %d равна %d\n", number, sum)
	

	
     
}