package main

import (
	"fmt"
	"net"
)

func main() {
	//主动连接服务器
	coon, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer coon.Close()

	coon.Write([]byte("are you ok?"))

}
