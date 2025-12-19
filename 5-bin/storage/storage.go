package storage

import "time"

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

//
//
//func (bin *Bin) SaveBin(name string) (bool, error) {
//	data, err := json.MarshalIndent(bin, "", "  ")
//	if err != nil {
//		return false, err
//	}
//	file, err := os.Create(name)
//	if err != nil {
//		fmt.Println(err)
//		return false, err
//	}
//	defer file.Close()
//	_, err = file.Write(data)
//	if err != nil {
//		return false, err
//	}
//	return true, nil
//}
//
//func (bin *Bin) LoadBin(filePath string) error {
//	data, err := os.ReadFile(filePath)
//	if err != nil {
//		return errors.New("could not read the files")
//	}
//	var result Bin
//	err = json.Unmarshal(data, &result)
//	if err != nil {
//		return errors.New("could not parse the files")
//	}
//	return nil
//
//}
