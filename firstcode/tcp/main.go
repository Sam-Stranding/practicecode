package main

import "os"

/*
tcp回音服务器
接收请求server，处理会话handleSession，
*/
func main() {
	exitChan := make(chan int)
	go server("127.0.0.1:7001", exitChan)
	code := <-exitChan
	//os.Exit(code int),表示 以执行状态码退出程序，0表示正常退出，非0表示异常退出，项目立即终止，函数延迟执行
	os.Exit(code)
}
