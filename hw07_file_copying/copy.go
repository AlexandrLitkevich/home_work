package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

//TODO Чтение из файла Done
//---TODO обработать EOF
//---TODO USE io.pipe
//TODO Запись в файл Done
//TODO Установка offset
//TODO Установка limit
//TODO Написание тестов

func Copy(fromPath, toPath string, offset, limit int64) error {
	file, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := make([]byte, 1000)

	r, err := io.ReadFull(file, buf)
	if err != nil {
		return err
	}
	fmt.Println("this result ReadFull", r)

	fileCopy, err := os.Create("test.txt")
	if err != nil {
		return err
	}

	defer fileCopy.Close()

	re, err := fileCopy.Write(buf)

	fmt.Println("this result Write", re)
	return nil
}
