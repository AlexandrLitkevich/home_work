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
//---TODO обработать EOF Done
//---TODO USE io.pipe maybe
//TODO Запись в файл Done
//TODO Установка offset
//TODO Установка limit
//TODO Написание тестов

// Tests
// TODO Копирование всего файла

//----CASES
// 1)Read all file and copy all file Done
//2) Use offset for read and copy file

func Copy(fromPath, toPath string, offset, limit int64) error {
	file, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	fmt.Println("fileInfo", fileInfo.Size())

	if offset >= fileInfo.Size() {
		return errors.New("invalid offset")
	}

	fileCopy, err := os.Create(toPath)
	if err != nil {
		return err
	}

	defer fileCopy.Close()

	if limit == 0 {
		_, err := io.Copy(fileCopy, file)
		if err != nil {
			return err
		}
		return nil
	}

	result := make([]byte, limit)

	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	_, err = file.Read(result)
	if err != nil {
		return err
	}

	_, err = fileCopy.Write(result)
	if err != nil {
		return err
	}

	return nil
}
