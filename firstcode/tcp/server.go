package main

import (
	"fmt"
	"net"
)

// 服务端，接收请求
func server(address string, exitChan chan int) {
	l, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}
	fmt.Println("listen address:", address)
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			exitChan <- 1
			continue
		}
		go handleSession(conn, exitChan)
	}
}
