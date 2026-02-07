package main

import (
	"fmt"
	"sync"
	"time"
)

/*
模拟电商秒杀
deduct -> 扣除
quantity -> 数量
*/

type Inventory struct {
	stock int
	rw    sync.RWMutex
}

func (v *Inventory) getStock() int {
	v.rw.RLock()
	defer v.rw.RUnlock()
	return v.stock
}
func (v *Inventory) deductStock(quantity int) bool {
	v.rw.Lock()
	defer v.rw.Unlock()
	if v.stock < quantity {
		fmt.Println("库存不足")
		return false
	}
	time.Sleep(time.Millisecond * 100)
	v.stock -= quantity
	return true
}

// 改为main
func cou() {
	inventory := Inventory{stock: 100}
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inventory.deductStock(1)
		}()
	}

	for i := 0; i < 20; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("剩余库存:", inventory.getStock())
	}

	wg.Wait()
	fmt.Println("最终剩余库存总量为：", inventory.getStock())
}
