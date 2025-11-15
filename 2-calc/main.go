package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	operation, number := getUserInput()
	//fmt.Println("Промежуток:", operation, number)
	result := work(operation, number)
	fmt.Println("Результат:", result)
}

func getUserInput() (string, []int) {

	var operation string
	fmt.Println("Введите операцию (AVG/SUM/MED)")
	fmt.Scan(&operation)

	fmt.Println("Введите числа через запятую (например: 1,2,3,4,5):")

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("Нет ввода")

	}

	input := scanner.Text()
	input = strings.TrimSpace(input)

	inputSplit := strings.Split(input, ",")

	var inputInt []int

	for _, x := range inputSplit {
		y, err := strconv.Atoi(x)
		if err != nil {
			fmt.Println(err)
		}
		inputInt = append(inputInt, y)
	}

	return operation, inputInt
}

func work(operation string, number []int) float64 {
	var res float64

	sort.Ints(number)

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
