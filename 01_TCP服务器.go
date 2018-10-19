package main

import (
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer listener.Close()

	con, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("err=", err1)
		return
	}

	buf := make([]byte, 1024)
	n, err2 := con.Read(buf)
	if err2 != nil {
		fmt.Println("err2=", err2)
		return
	}
	fmt.Println("buf=", string(buf[:n]))
	defer con.Close()

}
