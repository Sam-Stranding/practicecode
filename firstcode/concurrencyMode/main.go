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

/*
2.扇出模式，一个订单"流"分发给多个处理器，重点在"多个处理器"，让处理器并发执行
*/
func fanoutProcess(inputChan <-chan order, workID int, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range inputChan {
		switch workID {
		//workID: 1, 向商家，客户发送订单,如订单ID,价格
		case 1:
			fmt.Printf("订单ID为%d,价格为%.2f\n", order.ID, order.price)
		//workID: 2, 计算订单价格
		case 2:
			discounts := order.price * 0.8
			fmt.Printf("订单价格为: %.2f\n", discounts)
		//workID: 3, 写入数据库
		case 3:
			fmt.Println("订单已写入数据库")
		}
		time.Sleep(time.Millisecond * 100)
	}
}

/*
3.扇入模式，多个数据源合并 -> 多个生产者，由主goroutine收集订单
*/
func fanInProcess(outputCh chan<- order, ProducerID int) {
	defer fmt.Printf("生产者%d号生产订单")
	for i := 0; i < 3; i++ {
		Order := order{
			ID:       ProducerID*1000 + i + 1,
			UserID:   rand.Intn(1000),
			price:    rand.Float64() * 1000,
			status:   "new",
			createAt: time.Now(),
		}
		outputCh <- Order
		fmt.Printf("生产者%d生产订单%d号，价格为%.2f\n", ProducerID, Order.ID, Order.price)
		time.Sleep(time.Millisecond * 100)
	}
}

// 4.Pipeline模式L:订单处理流水线
func validationStage(input <-chan order) <-chan order {
	output := make(chan order, 10)
	go func() {
		//1.检验订单
		defer close(output)
		for order := range input {
			if order.price > 0 {
				order.status = "validated"
				fmt.Printf("订单已检验，订单ID为%d\n", order.ID)
				output <- order
			} else {
				fmt.Printf("订单校验失败，ID为%d\n", order.ID)
			}
		}
	}()
	return output
}
func paymentStage(input <-chan order) <-chan order {
	output := make(chan order, 10)
	go func() {
		defer close(output)
		//2.支付
		for order := range input {
			order.status = "paid"
			fmt.Printf("订单已支付，订单ID为%d\n", order.ID)
			output <- order
		}
	}()
	return output
}
func shippingStage(input <-chan order) <-chan order {
	output := make(chan order, 10)
	go func() {
		defer close(output)
		//3.配送
		for order := range input {
			order.status = "shipped"
			fmt.Printf("订单已配送，订单ID为%d\n", order.ID)
			output <- order
		}
	}()
	return output
}

func main() {
	//orderChan := make(chan order, 3)
	//var wg sync.WaitGroup

	////启动消费者(可以有多个消费者)
	//wg.Add(2)
	//go OrderConsume(orderChan, &wg)
	//go OrderConsume(orderChan, &wg)
	//
	//启动生产者
	//go OrderProduce(orderChan, 1)
	//wg.Wait()

	//扇出模式
	//fmt.Println("======扇出模式演示======")
	//inputChan := make(chan order, 10)
	//var fanwg sync.WaitGroup
	//
	//fanwg.Add(3)
	//for i := 0; i < 3; i++ {
	//	go fanoutProcess(inputChan, i+1, &fanwg)
	//}
	//
	////生成数据
	//go func() {
	//	for i := 0; i < 6; i++ {
	//		inputChan <- order{ID: i + 1, price: rand.Float64() * 1000}
	//		time.Sleep(time.Millisecond * 100)
	//	}
	//	//需要及时关闭channel,因为fanoutProcess()中使用了for range遍历channel，如果不关闭，goroutine会阻塞在range上
	//	close(inputChan)
	//}()
	//
	//fanwg.Wait()

	//扇入模式
	//fmt.Println("======扇入模式演示======")
	//fanInCh := make(chan order, 10)
	//
	////启动生产者
	//for i := 0; i < 3; i++ {
	//	go fanInProcess(fanInCh, i+1)
	//}
	//
	////处理订单
	//go func() {
	//	time.Sleep(time.Second * 2)
	//	close(fanInCh)
	//}()
	//
	//fmt.Println("收集到以下订单：")
	//for order := range fanInCh {
	//	fmt.Printf("订单ID为%d,价格为%.2f\n", order.ID, order.price)
	//}

	// pipeline模式
	fmt.Println("======pipeline模式演示======")
	pipeline := make(chan order, 10)

	validation := validationStage(pipeline)
	payment := paymentStage(validation)
	shipping := shippingStage(payment)

	//发送测试数据
	go func() {
		for i := 0; i < 3; i++ {
			Order := order{
				ID:       i + 1,
				UserID:   rand.Intn(1000),
				price:    rand.Float64() * 1000,
				status:   "new",
				createAt: time.Now(),
			}
			pipeline <- Order
		}
		close(pipeline)
	}()

	//打印结果
	for order := range shipping {
		fmt.Printf("%d号订单已完成，状态为%s\n", order.ID, order.status)
	}

}
