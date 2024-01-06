package main

import (
	"fmt"
	"os"
)

func main() {

	dir := os.Args[1]
	fmt.Println("this dir", dir)

	env, err := ReadDir(dir)
	if err != nil {
		panic(err)
	}

	fmt.Println(env)

	fmt.Println("this args[2]", os.Args[2])
}
