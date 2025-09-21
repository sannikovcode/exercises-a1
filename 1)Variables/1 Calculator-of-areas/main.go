package main

import "fmt"

func main() {
	var choice int
	fmt.Println(" Выберите фигуру: 1 - квадрат, 2 - прямоугольник, 3 - круг")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		var side float64
		fmt.Print("Введите длину стороны: ")
		fmt.Scanln(&side)
		area := side * side
		fmt.Printf("Площадь квадрата: %.2f\n", area)

	case 2:
		var length, width float64
		fmt.Print("Введите длину и ширину: ")
		fmt.Scanln(&length, &width)
		area := length * width
		fmt.Printf("Площадь прямоуольника: %.2f\n", area)

	case 3:
		var radius float64
		fmt.Print("Введите радиус: ")
		fmt.Scanln(&radius)
		area := 3.14159 * radius * radius
		fmt.Printf("Площадь круга: %.2f\n", area)

	}

}
