package main

import "fmt"

func main() {
	var fromCurrency, toCurrency string = "USD", "EUR"
	someValue1 := getUserInput()
	res := calc(someValue1, fromCurrency, toCurrency)

	fmt.Printf("Какой-то результат: %d", res)
}
func getUserInput() int {
	var inputVal1 int
	fmt.Println("Введите что-нибудь1")
	fmt.Scan(&inputVal1)

	return inputVal1
}

func calc(x int, y string, z string) int {
	result := 3
	return result
	// немного изменил
}
