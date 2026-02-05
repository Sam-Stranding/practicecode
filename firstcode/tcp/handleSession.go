package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// 处理session(会话)
func handleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("Session started:")

	//reader是用来读取conn的数据的一个类，具有缓冲区
	reader := bufio.NewReader(conn)
	for {
		/*
			ReadString()一直读取到输入中第一次出现delim，返回包含delim的字符串
			返回err != nil当且仅当返回的数据不以delim结束返回
			->只有err == nil时，str才是有效的
		*/
		str, err := reader.ReadString('\n')
		if err == nil {
			//去掉字符串末尾的回车
			str = strings.TrimSpace(str)
			//处理Telnet命令
			if !processTelnetCommand(str, exitChan) {
				conn.Close()
				break
			}
			//Echo逻辑，发什么数据，就返回什么数据
			conn.Write([]byte(str + "\r\n"))
		} else {
			fmt.Println("Session closed")
			conn.Close()
			break
		}
	}
}
