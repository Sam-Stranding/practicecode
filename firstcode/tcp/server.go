package main

import (
	"fmt"
	"net"
)

// 服务端，接收请求
func server(address string, exitChan chan int) {
	//监听address
	l, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		//出现错误，向通道写入1，os.Exit(1)表示服务器异常退出
		exitChan <- 1
	}
	fmt.Println("listen address:", address)
	defer l.Close()
	for {
		//接受请求,如果没有请求，将一直阻塞当前goroutine
		//从监听器接受一个新的客户端的连接，创建一个新的conn对象（代表该客户端连接），返回这个连接供服务器处理（读写）数据
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			exitChan <- 1
			continue
		}
		//创建一个goroutine处理会话
		go handleSession(conn, exitChan)
	}
}
