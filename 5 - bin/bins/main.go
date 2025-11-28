package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"time"

	"main.go/storage"
)

type Bin struct {
	Name      string
	Id        string
	CreatedAt time.Time
	Private   bool
}

func (acc *Bin) generateId(n int) {
	var validChars []rune

	for r := '0'; r <= '9'; r++ {
		validChars = append(validChars, r)
	}
	res := make([]rune, n)
	for i := 0; i < n; i++ {
		res[i] = validChars[rand.IntN(len(validChars))]
	}

	acc.Id = string(res)
}
func newBin(name, id string, private bool) (*Bin, error) {
	if name == "" {
		return nil, errors.New("invalid name")
	}

	newBin := &Bin{
		Name:      name,
		Id:        id,
		CreatedAt: time.Now(),
		Private:   private,
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

	storage.SaveBin(bin1, "bins.json")

	bin2, err := storage.LoadBin("bins.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bin2)
}
