package main

import (
	"fmt"
	"os"
)

func main() {
	dir := os.Args[1]
	fmt.Println("this dir", dir)
	cmds := os.Args[2:]
	fmt.Println("this dir", cmds)

	env, err := ReadDir(dir)
	if err != nil {
		panic(err)
	}

	fmt.Println(env)
}
