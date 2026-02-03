package main

import (
	"fmt"
	"time"
)

// GoroutineOne goroutine,
// 并发指同一时间不能运行多个任务；并行指同一时间可以运行多个任务
func GoroutineOne() {
	times := 1
	for {
		times++
		fmt.Printf("goruntine one: %d\n", times)
		//time.Sleep(time.Second),休眠1秒
		time.Sleep(time.Second)
	}
}
