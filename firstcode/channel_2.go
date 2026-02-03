package main

import (
	"fmt"
	"time"
)

// 循环接收，通过for data := range ch 循环接收数据
func Channel2() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
	}()
	for data := range ch {
		fmt.Println(data)
		if data == 2 {
			break
		}
	}
}
