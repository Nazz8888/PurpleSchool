package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"time"
)

type bin struct {
	name      string
	id        string
	createdAt time.Time
	private   bool
}

func (acc *bin) generateId(n int) {
	var validChars []rune

	for r := '0'; r <= '9'; r++ {
		validChars = append(validChars, r)
	}
	res := make([]rune, n)
	for i := 0; i < n; i++ {
		res[i] = validChars[rand.IntN(len(validChars))]
	}

	acc.id = string(res)
}
func newBin(name, id string, private bool) (*bin, error) {
	if name == "" {
		return nil, errors.New("invalid name")
	}

	newBin := &bin{
		name:      name,
		id:        id,
		createdAt: time.Now(),
		private:   private,
	}
	if id == "" {
		newBin.generateId(12)

	}
	return newBin, nil
}

func promptDataString(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
func promptDataBool(prompt string) bool {
	fmt.Print(prompt)
	var res bool
	fmt.Scanln(&res)
	return res
}

func main() {
	name := promptDataString("Введите имя: ")
	private := promptDataBool("Бин приватный? ")
	id := promptDataString("Введите id: ")

	bin1, err := newBin(name, id, private)
	if err != nil {
		fmt.Println("неверный формат ввода")
		return
	}

	BinList := make(map[string]bin)
	fmt.Println(bin1)
	fmt.Println(BinList)

}
