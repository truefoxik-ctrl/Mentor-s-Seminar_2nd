package main

import "fmt"

func main() {
	var n int

	fmt.Println("Введите количество чисел")
	fmt.Scan(&n)

	var firstNum int
	fmt.Println("Введите число")
	fmt.Scan(&firstNum)
	var max_n int = firstNum
	var min_n int = firstNum

	for i := 1; i < n; i++ {
		var x int
		fmt.Println("Введите число")
		fmt.Scan(&x)
		if x > max_n {
			max_n = x
		}

		if x < min_n {
			min_n = x
		}
	}

	fmt.Println("Максимальное значение среди введенных чисел:", max_n)
	fmt.Println("Минимальное значение среди введенных чисел:", min_n)

}
