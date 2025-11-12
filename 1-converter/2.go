package main

import "fmt"

func main() {
	var userName string
	var userAge int
	userName, userAge = getUserInput()

	fmt.Printf("Привет, %s! Тебе %d лет.", userName, userAge)

}
func getUserInput() (string, int) {
	var userName string
	var userAge int
	fmt.Println("Введите своё имя")
	fmt.Scan(&userName)
	if len(userName) > 50 {
		fmt.Println("Ошибка ввода имени")

	}
	fmt.Println("Введите свой возраст")
	fmt.Scan(&userAge)
	if (userAge < 1) || (userAge > 120) {
		fmt.Println("Ошибка ввода возраста")

	}

	return userName, userAge
}
