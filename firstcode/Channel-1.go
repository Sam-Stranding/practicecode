package main

import (
	"fmt"
	"os"
	"sync"
)

/*
	Channel是通道，用于在多个goroutine之间进行通信
	同一时间只能有一个goroutine访问channel，进行数据发送和接收

	并发同步 指控制多个并发执行单元对共享资源的访问顺序，保证程序顺序执行
要求3点：
	1.互斥，同一时间只能有一个goroutine访问资源
	2.可见性，一个goroutine对资源的修改，其余goroutine是可见的
	3。有序性，控制操作的执行顺序
*/
//<-ch 忽略接收数据，仅用于并发同步
func main() {
	//1.channel通道，通过<- ch 忽略接收数据，仅用于并发同步
	ch := make(chan int)
	go func() {
		fmt.Println("one")
		ch <- 1
		fmt.Println("two")
	}()
	fmt.Println("three")
	<-ch
	fmt.Println("four")

	//2.互斥锁
	/*
		sync.WaitGroup(简称wg)是GO的一个标准库，用于协调多个goroutine执行完成的一个同步原语，
		让主goroutine等待后台的goroutine执行完成。
		主要是三板斧：
		1.wg.Add(delta int),表示后台有delta个goroutine
		2.wg.Done(),表示后台的goroutine执行完成
		3.wg.Wait(),阻塞当前goroutine，当内部计数器归零时，wg.Wait()返回
	*/
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		//1.用err捕捉可能发生的panic，最后再释放锁
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic:", err)
			}
			mu.Unlock()
		}()
		//2.上锁
		mu.Lock()
		//3.用data打开文件，注意，此时data是*os.File类型,如果使用fmt.Printf(data),
		data, err2 := os.Open("firstcode/channel-1")
		if err2 != nil {
			fmt.Printf("open file error: %v\n", err2)
			return
		}
		defer data.Close()
		//4.读data
		buf := make([]byte, 1024)
		n, err3 := data.Read(buf)
		if err3 != nil {
			fmt.Printf("read file error: %v\n", err3)
			return
		}
		fmt.Println(string(buf[:n]))
	}()
	wg.Wait()
}
