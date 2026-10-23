package main

import "fmt"

func main() {
	var x, y int
	var action string
	fmt.Scan(&x, &y)
	fmt.Scan(&action)

	switch action {
	case "+":
		fmt.Print(x, "+", y, "=", x+y)
	case "-":
		fmt.Print(x, "-", y, "=", x-y)
	case "*":
		fmt.Print(x, "*", y, "=", x*y)
	case "/":
		if y != 0 {
			fmt.Print(x, "/", y, "=", x/y)
		} else {
			fmt.Print("Деление на 0 невозможно")
		}
	}
}
