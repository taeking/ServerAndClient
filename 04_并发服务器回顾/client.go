package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	con, err1 := net.Dial("tcp", "127.0.0.1:8000")
	if err1 != nil {
		fmt.Println("err1=", err1)
	}
	go func() {
		str := make([]byte, 1024)
		for {
			n, err := con.Read(str)
			if err != nil {
				fmt.Println("err=", err)
			}
			fmt.Println(string(str[:n]))
		}

	}()

	buf := make([]byte, 1024)
	for {

		n, err2 := os.Stdin.Read(buf)
		if err2 != nil {
			fmt.Println("err2=", err2)
		}

		con.Write(buf[:n])
	}

}
