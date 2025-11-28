package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Bin struct {
	Name      string
	Id        string
	CreatedAt time.Time
	Private   bool
}

func SaveBin(bin *Bin, name string) (bool, error) {
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

func LoadBin(filePath string) (*Bin, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("could not read the file")
	}
	var result Bin
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.New("could not parse the file")
	}
	return &result, nil

}
