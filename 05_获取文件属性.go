package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("message error")

	}
	filename := list[1]

	fileinfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println("err=", err)

	}
	fmt.Println("name=", fileinfo.Name())
	fmt.Println("size=", fileinfo.Size())

}
