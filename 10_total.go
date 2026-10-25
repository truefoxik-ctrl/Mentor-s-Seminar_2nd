package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите количество чисел")
	fmt.Scan(&n)

	fmt.Println("Введите первое число")
	var firstNum int
	fmt.Scan(&firstNum)

	var max_n int = firstNum
	var min_n int = firstNum
	var sum_n int = firstNum
	var even_n int
	var uneven_n int

	if firstNum%2 == 0 {
		even_n += 1
	} else {
		uneven_n += 1
	}

	for i := 1; i < n; i++ {
		var x int
		fmt.Println("Введите число")
		fmt.Scan(&x)

		sum_n += x

		if x > max_n {
			max_n = x
		}
		if x < min_n {
			min_n = x
		}
		if x%2 == 0 {
			even_n += 1
		} else {
			uneven_n += 1
		}

	}
	var avg_n float64 = float64(sum_n) / float64(n)

	fmt.Println("Сумма всех", n, "чисел равна", sum_n)
	fmt.Println("Среднее всех чисел равно", avg_n)
	fmt.Println("Максимальное из введенных чисел равно", max_n)
	fmt.Println("Минимальное из введенных чисел равно", min_n)
	fmt.Println("Количество четных чисел в списке:", even_n)
	fmt.Println("Количество нечетных чисел в списке:", uneven_n)

}
