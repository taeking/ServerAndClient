package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	con, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err=", err)

	}
	defer con.Close()

	go func() {

		buf := make([]byte, 1024)
		for {
			n, err2 := con.Read(buf)
			if err2 != nil {
				fmt.Println("err2=", err2)
				return
			}
			fmt.Println(string(buf[:n]))
		}

	}()
	str := make([]byte, 1024)
	for {
		n, err2 := os.Stdin.Read(str)
		if err2 != nil {
			fmt.Println("err2=", err2)
			return
		}
		con.Write(str[:n])

	}

}
