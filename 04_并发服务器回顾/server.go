package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleCon(con net.Conn) {
	addr := con.RemoteAddr().String()
	fmt.Println("addr connect succeful", addr)
	buf := make([]byte, 1024)
	for {
		n, err := con.Read(buf)
		if err != nil {
			fmt.Println("err=", err)
		}
		fmt.Println(string(buf[:n]))
		con.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}

func main() {
	listener, err1 := net.Listen("tcp", "127.0.0.1:8000")
	if err1 != nil {
		fmt.Println("err1=", err1)
	}
	for {
		con, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("err2=", err2)
		}
		go HandleCon(con)
	}

}
