package entity

import (
	"companyAgancies/constant"
	"fmt"
	"os"
	"strings"
)

type Storage struct {
	Path string
}

func New(path string) Storage {
	s := Storage{}
	s.SetPath(path)
	return s
}

func (storage *Storage) SetPath(path string) {
	if storage.Path == "" {
		path = constant.DefaultPath
	}
	storage.Path = path
}

func (storage *Storage) StoreIntoFile(data []byte) {
	var file *os.File

	file, fErr := os.OpenFile(storage.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if fErr != nil {
		fmt.Println("there is an error: ", fErr)
	}
	data = append(data, []byte("\n")...)

	_, wErr := file.Write(data)
	if wErr != nil {
		fmt.Println("error in writing", wErr)
	}
	file.Close()
}

func (storage *Storage) ReadFromFile() []string {
	fileData, err := os.ReadFile(storage.Path)
	if err != nil {
		fmt.Println("error: ", err)
	}

	var dataStr = string(fileData)
	return strings.Split(dataStr, "\n")

}
