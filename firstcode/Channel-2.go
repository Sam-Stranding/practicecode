package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		ch := make(chan int)
		for i := 0; i < 3; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		for data := range ch {
			fmt.Println(data)
			if data == 0 {
				break
			}
		}
	}()
}
