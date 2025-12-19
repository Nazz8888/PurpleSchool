package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"time"
)

type Bin struct {
	Name      string
	Id        string
	CreatedAt time.Time
	Private   bool
}
type WorkBin interface {
	SaveBin(name string) (bool, error)
	LoadBin(filePath string) error
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

func (bin *Bin) SaveBin(name string) (bool, error) {
	data, err := json.MarshalIndent(bin, "", "  ")
	if err != nil {
		return false, err
	}
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (bin *Bin) LoadBin(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return errors.New("could not read the files")
	}
	var result Bin
	err = json.Unmarshal(data, &result)
	if err != nil {
		return errors.New("could not parse the files")
	}
	return nil

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

	name := &Bin{}
	name.Name = promptDataString("Введите имя: ")
	name.Private = promptDataBool("Бин приватный? ")
	name.generateId(12)
	name.CreatedAt = time.Now()

	name.SaveBin("bins.json")

	//
	//	storage.WorkBin.SaveBin("bins.json")
	//fmt.Println(name)

	//name := promptDataString("Введите имя: ")
	//private := promptDataBool("Бин приватный? ")
	//id := promptDataString("Введите id: ")
	//
	//var bin1 Bin
	//bin1, err = newBin(name, id, private)
	//if err != nil {
	//	fmt.Println("неверный формат ввода")
	//	return
	//}
	//
	//storage.WorkBin.SaveBin(bin1, bin1.Name)

	//storage.SaveBin(bin1, "bins.json")
	//
	//bin2, err := storage.LoadBin("bins.json")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(bin2)
}
