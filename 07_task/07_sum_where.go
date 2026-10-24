package main

import "fmt"

func main() {
	var x int
	fmt.Println("Введите число N")
	fmt.Scan(&x)

	var sum_x int = 0

	for i := 1; i <= x; i++ {
		if i%6 == 0 {
			sum_x += i
		}
	}

	fmt.Println("Сумма четных чисел, от 1 до", x, ", которые делятся на 3, равна", sum_x)

}
