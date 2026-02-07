package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
原子操作
一般使用atomic
*/

type counter struct {
	count int64
}

func (c *counter) inc() {
	//time.Sleep(time.Millisecond * 10)
	atomic.AddInt64(&c.count, 1)
}
func (c *counter) dec() {
	//time.Sleep(time.Millisecond * 20)
	atomic.AddInt64(&c.count, -1)
}
func (c *counter) get() int64 {
	return atomic.LoadInt64(&c.count)
}

func main() {
	c := counter{}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.inc()
		}()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.dec()
		}()
	}
	for i := 0; i < 20; i++ {
		fmt.Println("点赞数：", c.get())
	}
	wg.Wait()
	fmt.Println("最终点赞数：", c.get())
}
