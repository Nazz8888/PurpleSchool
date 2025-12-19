package file

import (
	"encoding/json"
	"errors"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("could not read the file")
	}

	if json.Valid(data) {
		return data, nil
	} else {
		return nil, errors.New("invalid file")
	}
}

func WriteFile() {

}
