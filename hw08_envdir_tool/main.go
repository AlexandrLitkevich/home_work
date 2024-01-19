package main

import (
	"os"
)

func main() {
	dir := os.Args[1]
	cmds := os.Args[2:]

	envs, err := ReadDir(dir)
	if err != nil {
		panic(err)
	}
	result := RunCmd(cmds, envs)
	os.Exit(result)
}
