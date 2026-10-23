package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if x > 0 {
		fmt.Println("Число", x, "положительное")
	} else if x < 0 {
		fmt.Println("Число", x, "отрицательное")
	} else {
		fmt.Println("0 не является ни положительным, ни отрицательным числом")
	}

	if x%2 == 0 {
		fmt.Println("Число", x, "четное")
	} else {
		fmt.Println("Число", x, "нечетное")
	}

	if x%3 == 0 {
		fmt.Println("Число", x, "делится без остатка на 3")
	} else {
		fmt.Println("Число", x, "не делится без остатка на 3")
	}

	if x%5 == 0 {
		fmt.Println("Число", x, "делится без остатка на 5")
	} else {
		fmt.Println("Число", x, "не делится без остатка на 5")
	}

}
