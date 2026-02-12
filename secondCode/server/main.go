package main

import (
	"fmt"
	"net/http"
)

// 理解net/http包
func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("req: ping")
		w.Write([]byte("pong")) //给客户端 写一个字符串 pong
		/*
			详细讲一下这个w.Write(),为什么不能w.Write("pong")?
			在源码里面，Write([]byte) (int, error)，接收的是[]byte，而不是string
			本质在于 底层I/O传输的是字节，不是字符串这种高级的数据结构
		*/
	})
	/*
			传入nil,表示使用默认的http.ServeMux:
				需要传入 handler ,即实现ServeHTTP方法的任意类型，
				在ServeHTTP中，有 if handler == nil{handler == DefaultServeMux}
				DefaultServeMux = &defaultServeMux
				var defaultServeMux ServeMux
			DefaultServeMux 是一个全局路由器，可以使用http.HandleFunc()/http.Handle()注册路由
			如果传入nil,即使用默认的路由器。
		此方法适用于简单快捷，小脚本
	*/
	http.ListenAndServe(":8080", nil)
}
