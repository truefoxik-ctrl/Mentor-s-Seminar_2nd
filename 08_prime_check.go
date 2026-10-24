package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Println("Введите число")
	fmt.Scan(&n)

	if isPrime(n) {
		fmt.Println(n, "является простым числом")
	} else {
		fmt.Println(n, "не является простым числом")
	}
}
