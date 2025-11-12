package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	const usd2rub = 81.29
	const eur2usd = 1.16
	const eur2rub = 94.19

	inputVal1, inputVal2, inputVal3 := getUserInput()
	inputVal4 := float64(inputVal2)
	var result float64

	switch inputVal1 {
	case "rub":
		switch inputVal3 {
		case "eur":
			result = inputVal4 / eur2rub
		case "usd":
			result = inputVal4 / usd2rub
		}

	case "eur":
		switch inputVal3 {
		case "rub":
			result = inputVal4 * eur2rub
		case "usd":
			result = inputVal4 / eur2usd
		}

	case "usd":
		switch inputVal3 {
		case "rub":
			result = inputVal4 * usd2rub
		case "eur":
			result = inputVal4 * eur2usd
		}

	}

	fmt.Println("результат", result)
}

func getUserInput() (string, int, string) {

	var inputVal1, inputVal3 string
	var inputVal2 int
	for {
		fmt.Println("Введите исходную валюту (RUB/USD/EUR)")
		var x string
		var err error
		fmt.Scan(&x)
		x, err = valCheckStr(x)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			inputVal1 = x
			break
		}
	}

	for {
		fmt.Println("Введите число")
		var y int
		var err error
		fmt.Scan(&y)
		y, err = valCheckInt(y)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			inputVal2 = y
			break
		}

	}

	for {
		fmt.Println("Введите целевую валюту (RUB/USD/EUR)")
		var z string
		var err error
		fmt.Scan(&z)
		z, err = valCheckStr(z)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			inputVal3 = z
			break
		}
	}
	return inputVal1, inputVal2, inputVal3
}

func valCheckStr(i string) (string, error) {
	x := strings.ToLower(i)
	switch x {
	case "rub":
		return x, nil
	case "usd":
		return x, nil
	case "eur":
		return x, nil
	default:
		return "", errors.New("Не задана валюта для расчета")
	}
}

func valCheckInt(x int) (int, error) {
	if x <= 0 {
		return x, errors.New("Не задано число для расчета")
	} else {
		return x, nil
	}
}
