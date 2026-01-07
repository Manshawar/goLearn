package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	// 情况1：有接收者在等待
	// go func() {
	// 	<-messages // 接收者在等待
	// }()
	go func() {
		// time.Sleep(time.Second * 1)
		messages <- "hello"
	}()
	// msg := "hi"
	// select {
	// case messages <- msg: // ✅ 会执行，因为有接收者
	// 	fmt.Println("sent message", msg)
	// default:
	// 	fmt.Println("no message sent")
	// }
	// 问题：为什么 select 中的接收一直走 default？
	// 答案：这是 goroutine 调度时序问题
	// select 执行时，goroutine 可能还没发送消息，通道为空
	// 因为有 default，不会阻塞，直接执行 default
	// 解决：去掉 default 让 select 阻塞等待，或使用缓冲通道
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	// 问题：select 一般为接收的，写入可能执行吗？
	// 答案：可以！select 不仅可以用于接收，也可以用于发送
	// select 是多路复用，可以处理发送和接收
	// default 用于非阻塞操作，如果所有 case 都不能立即执行，就执行 default
	// 因为这是无缓冲区通道，并且也没有接收者，因此，default 会执行
	// 有了 default 就直接走默认了
	// 情况2：使用缓冲通道
	buffered := make(chan string, 1) // 容量为 1

	select {
	case buffered <- "hello": // ✅ 会执行，因为缓冲区有空间
		fmt.Println("sent message")
	default:
		fmt.Println("no message sent")
	}
}
