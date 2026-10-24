package main

import "fmt"

func main() {
	var n int

	fmt.Println("Введите количество чисел")
	fmt.Scan(&n)
	var null_counter int = 0
	var positive_counter int = 0
	var negative_counter int = 0

	for i := 0; i < n; i++ {
		var x int
		fmt.Println("Введите число")
		fmt.Scan(&x)
		if x > 0 {
			positive_counter += 1
		} else if x < 0 {
			negative_counter += 1
		} else {
			null_counter += 1
		}
	}
	fmt.Println("Количество нулей:", null_counter)
	fmt.Println("Количество положительных чисел:", positive_counter)
	fmt.Println("Количество отрицательных чисел:", negative_counter)
}
