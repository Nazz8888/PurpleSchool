package files

import (
	"errors"
	"os"
	"path/filepath"
)

func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("could not read the files")
	}
	if filepath.Ext(path) == ".json" {

		return data, nil
	} else {
		return nil, errors.New("invalid files")
	}
}

func WriteFile() {

}
