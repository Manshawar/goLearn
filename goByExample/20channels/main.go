package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second * 2)
	fmt.Println("done")
	done <- true
}
func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(ping <-chan string, pongs chan<- string) {
	msg := <-ping
	pongs <- msg
}
func main() {
	message := make(chan string)

	go func() { message <- "ping" }()
	// 问题：通道不能直接打印吗？<-message 拿值只能使用一次吗？
	// 答案：
	// 1. 可以打印 <-message，这是从通道接收值并打印
	// 2. <- 可以多次使用，但每次只取一个值，取走后值就从通道中移除了
	// 3. 通道中有多个值时，可以多次取；通道为空时，再次取会阻塞
	fmt.Println("通道类型/", <-message) // 输出: 通道类型: chan string
	// msg := <-message  // 如果再次取，通道已经空了，会阻塞
	message1 := make(chan string, 2)
	message1 <- "buffered"
	message1 <- "channel"
	fmt.Println(<-message1)
	fmt.Println(<-message1)
	done := make(chan bool, 1)
	go worker(done)
	<-done
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
	// fmt.Println(msg)
}
