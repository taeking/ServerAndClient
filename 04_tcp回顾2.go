package main

import (
	"fmt"
	"net"
)

func main() {
	con, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err=", err)
		return

	}
	//buf := make([]byte, 1024)
	con.Write([]byte("are you ok"))

}
