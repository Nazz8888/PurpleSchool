package main

import "fmt"

func main() {

	operation, number := getUserInput()
	//fmt.Println("Промежуток:", operation, number)
	result := work(operation, number)
	fmt.Println("Результат:", result)
}

func getUserInput() (string, []int) {
	var operation string
	var number []int
	var num int

	fmt.Println("Введите операцию (AVG/SUM/MED)")
	fmt.Scan(&operation)

	for {
		fmt.Println("Введите число")
		fmt.Scan(&num)

		if num == 0 {
			break
		}
		number = append(number, num)

	}
	return operation, number
}

func work(operation string, number []int) float64 {
	var res float64
	switch operation {
	case "AVG":
		var sum float64
		var count int
		for _, y := range number {
			sum += float64(y)
			count++

		}
		res = float64(sum / float64(count))

	case "SUM":
		var sum int
		for _, y := range number {
			sum += y
		}
		res = float64(sum)

	case "MED":
		var med float64
		if len(number)%2 != 0 {
			med = float64(number[len(number)/2])
		}
		if len(number)%2 == 0 {
			med = float64(number[(len(number)/2)-1]+number[(len(number)/2)]) / 2
		}
		res = med

	}
	return res
}
