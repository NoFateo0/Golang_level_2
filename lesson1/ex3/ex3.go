package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type MyError struct {
	errorTime time.Time
	path      string
	errorStr  string
}

func (m MyError) Error() string {
	return fmt.Sprintf("%s\nНевозможно создать файл: %s\nПо причине: %s",
		m.errorTime, m.path, m.errorStr)
}

func CreateFile(dirPath, fileName string) error {
	fullPath := filepath.Join(dirPath, fileName)
	file, err := os.Create(fullPath)
	if err != nil {
		return MyError{
			path:      fullPath,
			errorStr:  err.Error(),
			errorTime: time.Now(),
		}
	}

	defer file.Close()

	return nil
}

func main() {
	for i := 0; i < 30; i++ {
		iStr := strconv.Itoa(i)
		if err := CreateFile("E:\\Sources\\Golang_GB_2\\lesson1\\test", "test"+iStr+".txt"); err != nil {
			fmt.Println(err.Error())
		}
	}
}
