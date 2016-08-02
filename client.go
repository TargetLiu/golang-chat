package main

import (
	"fmt"
	"net"
	"strings"
)

func startClient(port string) {
	//创建TCP请求地址
	addr, err := net.ResolveTCPAddr("tcp", port)
	checkErr(err)

	//连接服务端
	conn, err := net.DialTCP("tcp", nil, addr)
	checkErr(err)

	//读取欢迎信息
	data := make([]byte, 1024)
	conn.Read(data)
	fmt.Println(string(data))

	//输入昵称
	fmt.Print("Please input your nickname:")
	nickname := ScanLine()
	fmt.Println("Hello " + nickname)
	conn.Write([]byte("hello|" + nickname))

	//开启协程处理消息
	go handle(conn, nickname)

	//发送消息
	for {
		message := ScanLine()
		conn.Write([]byte("say|" + nickname + "|" + message))
	}
}

func handle(conn net.Conn, nickname string) {
	for {
		data := make([]byte, 1024)
		//读取消息
		_, err := conn.Read(data)
		checkErr(err)
		//屏蔽自身发送的聊天内容
		if strings.Contains(string(data), "["+nickname+"]: ") == false {
			fmt.Println(string(data))
		}
	}
}
