package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConnect(con net.Conn) {
	defer con.Close()
	//获取客户端网络信息

	addr := con.RemoteAddr().String()
	fmt.Println("addr connect succeful")
	buf := make([]byte, 2048)
	for {
		n, err := con.Read(buf)
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		fmt.Println(len(string(buf[:n-2])))

		if "exit" == string(buf[:n-2]) {

			fmt.Println("exit", addr)
			return
		}

		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))

		con.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer listener.Close()

	for {
		con, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println("err=", err1)
			return
		}

		go HandleConnect(con)
	}
}
