package main

import (
	"fmt"
	"strings"
)

//Telnet命令处理逻辑
/*
@close 退出当前连接会话
@shutdown 终止服务器运行
*/
func processTelnetCommand(str string, exitChan chan int) bool {
	//1."@close"指令表示终止本次会话->关闭conn
	//2."@shutdown"指令表示终止服务器运行->退出程序
	//strings.HasPrefix(s, prefix string)bool{} 判断s是否以prefix开头
	if strings.HasPrefix(str, "@close") {
		fmt.Println("Session closed")
		return false
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("Server shutdown")
		//向通道写入0，表示关闭服务器
		exitChan <- 0
		return false
	}
	fmt.Println(str)
	return true
}
