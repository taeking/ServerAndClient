package main

import (
	"fmt"
	"net"
)

type Client struct {
	C    chan string //用户发送数据的管道
	Name string      //用户名
	Addr string      //网络地址
}

func HandleConnec(con net.Conn) { //处理用户链接
	//获取客户端网络地址
	cliAddr := con.RemoteAddr().String()
	//创建一个结构体
	cli := Client{make(chan string), cliAddr, cliAddr}
	onlineMap[cliAddr] = cli
	//新开一个协程专门给当前客户端发送信息
	go Writemsgtoclient(cli, con)
	message <- "[" + cli.Addr + "]" + cli.Name + ":login"
	for {

	}

}
func Writemsgtoclient(cli Client, con net.Conn) {
	for msg := range cli.C {
		con.Write([]byte(msg + "\n"))
	}

}

//保存在线用户
var onlineMap map[string]Client
var message = make(chan string)

func Manager() {
	//给map分配空间
	onlineMap = make(map[string]Client)

	for {
		msg := <-message //没有消息这里会阻塞
		for _, cli := range onlineMap {

			cli.C <- msg
		}
	}

}

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer listener.Close()

	//新开一个协程，转发消息，只要有消息来了，遍历map给每个map成员都发送消息
	go Manager()
	//主协程，循环阻塞用户链接
	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Println("err=", err)
			continue

		}
		go HandleConnec(con)

	}

}
