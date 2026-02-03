package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

/*
RPC 模拟
需要注意使用WaitGroup，因为 主协程可能提前结束，导致部分goroutine被直接中止
模拟客户端与服务器的方法中不需要写wg.Done(),因为它俩只是方法，在主协程中的go func(){}()中使用就可以
wg可以作为参数传入方法，但必须传入指针，即&wg。
启动模拟超时，会出现channel阻塞，ch <- "ok"无接收
*/

// RPCClient 模拟客户端发送请求，接收数据，如果超时，触发error
func RPCClient(ch chan string, req string) (string, error) {
	//1.先发送请求
	ch <- req
	//2.接收数据，超时处理
	select {
	case ack := <-ch:

		fmt.Println("client received:", ack)
		return ack, nil
	//超时处理
	case <-time.After(time.Second):
		return "", errors.New("time out")
	}

}

// RPCServer 模拟服务器，接收请求，打印数据，返回数据
func RPCServer(ch chan string) {
	data := <-ch
	fmt.Println("Server received:", data)
	//模拟超时处理
	//time.Sleep(time.Second * 2)
	ch <- "ok"
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		ack, err := RPCClient(ch, "hi")
		if err != nil {
			fmt.Println("client error:", err)
		} else {
			fmt.Println("client content:", ack)
		}
	}()
	RPCServer(ch)
	wg.Wait()
}
