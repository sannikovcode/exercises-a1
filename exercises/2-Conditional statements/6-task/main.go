package main

import "fmt"


func main() {
	var age int
    
	fmt.Println("Введите ваш возраст:")
	fmt.Scan(&age)

	if age < 0{
		fmt.Println("Ошибка: возраст не может быть отрицательным!")
	} else if age < 18 {
		fmt.Println("Вы несовершеннлетний!")
	} else {
		fmt.Println("Вы совершеннолетний!")
	}
}