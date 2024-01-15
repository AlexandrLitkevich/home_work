package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	envs := make(Environment)
	for _, file := range files {
		name := strings.TrimRight(file.Name(), "=")
		currentFile, err := os.Open(path.Join(dir, name))
		if err != nil {
			currentFile.Close()
			fmt.Println("fail to open file", name)
		}

		envs[name] = ReadCurrentFile(currentFile)

		currentFile.Close()
	}
	return envs, nil
}

func ReadCurrentFile(in io.Reader) EnvValue {
	scanner := bufio.NewScanner(in)

	if scanner.Scan() {
		value := string(bytes.ReplaceAll([]byte(scanner.Text()), []byte{0x00}, []byte("\n")))

		return EnvValue{
			Value:      value,
			NeedRemove: false,
		}
	}

	return EnvValue{NeedRemove: true}
}
