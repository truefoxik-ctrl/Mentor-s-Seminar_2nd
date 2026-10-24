package main

import "fmt"

func printTable(x int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(x, "*", i, "=", x*i)
	}
}

func main() {
	var x int
	fmt.Println("Введите натуральное число от 1 до 10")
	fmt.Scan(&x)

	printTable(x)
}
