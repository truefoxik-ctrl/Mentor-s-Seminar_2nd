package main

import "fmt"

func main() {
	var x int
	fmt.Println("Введите число N")
	fmt.Scan(&x)

	var sum_x int = 0

	for i := 1; i <= x; i++ {
		sum_x += i
	}
	var avg_x float64 = float64(sum_x) / float64(x)

	fmt.Println("Сумма всех чисел от 1 до", x, "равна", sum_x)
	fmt.Println("Среднее всех чисел от 1 до", x, "равно", avg_x)

}
