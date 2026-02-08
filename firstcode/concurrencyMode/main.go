package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*

 */

type order struct {
	ID       int
	UserID   int
	price    float64 //价格
	status   string  //订单状态
	createAt time.Time
}

//1.生产者，消费者模型

func OrderProduce(orderChan chan<- order, orderNum int) {
	defer close(orderChan)
	for i := 1; i <= orderNum; i++ {
		Order := order{
			ID:       i,
			UserID:   rand.Intn(1000),
			price:    rand.Float64() * 1000,
			status:   "consignment", //consignment:寄售
			createAt: time.Now(),
		}
		//用随机数 * 时间，随机数需要使用time.Duration()时间间隔
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		orderChan <- Order
		fmt.Printf("订单已生产，ID为%d，价格为%.2f\n", i, Order.price) //模拟生产间隔
	}

}

func OrderConsume(orderChan <-chan order, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range orderChan {
		order.status = "completed" //completed: 完成
		fmt.Printf("处理订单：ID为%d，UserID为%d，价格为%.2f\n", order.ID, order.UserID, order.price)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	orderChan := make(chan order, 5)
	var wg sync.WaitGroup

	//启动消费者(可以有多个消费者)
	wg.Add(2)
	go OrderConsume(orderChan, &wg)
	go OrderConsume(orderChan, &wg)

	//启动生产者
	go OrderProduce(orderChan, 8)
	wg.Wait()
}
