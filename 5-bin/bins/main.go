package main

import (
	"5-bin/files"
	"5-bin/storage"
	"fmt"
	"math/rand/v2"
	"time"
)

type Bin struct {
	Name      string
	Id        string
	CreatedAt time.Time
	Private   bool
}

type Create interface {
	newBin(name, id string, private bool) (*Bin, error)
	promptDataString(prompt string) string
	promptDataBool(prompt string) bool
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

//устраняем дичь с коммитами
//func newBin(name, id string, private bool) (*Bin, error) {
//	if name == "" {
//		return nil, errors.New("invalid name")
//	}
//
//	newBin := &Bin{
//		Name:      name,
//		Id:        id,
//		CreatedAt: time.Now(),
//		Private:   private,
//	}
//	if id == "" {
//		newBin.generateId(12)
//
//	}
//	return newBin, nil
//}

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

	bin := &Bin{}
	bin.Name = promptDataString("Введите имя: ")
	bin.Private = promptDataBool("Бин приватный? ")
	bin.generateId(12)
	bin.CreatedAt = time.Now()

	var storage WorkBin = bin
	processBin(storage)

}
func processBin(storage WorkBin) {
	storage.SaveBin("bins.json")
	storage.LoadBin("bins.json")
}
