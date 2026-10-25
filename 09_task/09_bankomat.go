package main

import "fmt"

func main() {
	fmt.Println("Введите баланс")
	var balance float64
	fmt.Scan(&balance)
	fmt.Println("Введите команду")
	var action int
	fmt.Scan(&action)

	switch action {

	case 1:
		fmt.Println("Текущий баланс:", balance)

	case 2:
		fmt.Println("Введите сумму, на которую нужно пополнить баланс")
		var balance_plus float64
		fmt.Scan(&balance_plus)
		balance += balance_plus
		fmt.Println("Текущий баланс:", balance)

	case 3:
		fmt.Println("Введите сумму для снятия")
		var balance_minus float64
		fmt.Scan(&balance_minus)
		if balance_minus <= balance {
			balance -= balance_minus
			fmt.Println("Текущий баланс:", balance)
		} else {
			fmt.Println("Недостаточно средств")
		}
	}
}
