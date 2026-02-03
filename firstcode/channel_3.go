package main

import (
	"fmt"
	"time"
)

func Printer(c chan int) {
	defer func() {
		//recover()用于捕获panic，使程序继续执行，只能在defer中
		if i := recover(); i != nil {
			fmt.Println("Printer协程发生 panic: ", i)
		}
	}()
	for data := range c {
		fmt.Println(data)
	}
}

// Channel3 改为main
func Channel3() {
	//带缓冲大小的通道
	c := make(chan int, 10)
	go Printer(c)
	func() {
		defer close(c)
		defer func() {
			if i := recover(); i != nil {
				fmt.Println("main 协程发生 panic: ", i)
			}
		}()
		for i := 1; i <= 10; i++ {
			c <- i
		}
	}()
	//防止主协程退出过早，导致Printer协程未执行就退出
	time.Sleep(time.Second)
}
