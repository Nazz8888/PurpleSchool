package main

import "fmt"

func main() {

	var m map[string]string
	m = make(map[string]string)
	for {
		fmt.Println("Введите операцию")
		fmt.Println("1 Посмотреть закладки")
		fmt.Println("2 Добавить закладку")
		fmt.Println("3 Удалить закладку")
		fmt.Println("4 Выход")

		var operation int
		fmt.Scan(&operation)
		if operation == 1 {
			fmt.Println(m)
		}
		if operation == 2 {

			var name, value string
			fmt.Println("Введите название для добавления")
			fmt.Scan(&name)
			fmt.Println("Введите значение для добавления")
			fmt.Scan(&value)
			m[name] = value
		}

		if operation == 3 {
			var name string
			fmt.Println("Введите название для удаления")
			fmt.Scan(&name)
			delete(m, name)
		}

		if operation == 4 {
			break
		}

	}
}
