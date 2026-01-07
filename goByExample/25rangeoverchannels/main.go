package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// 问题：close 起什么作用？可以用通俗的话来解析吗？
	// 答案：
	// close() = 在通道上贴"已关闭，不再接收新数据"的标签
	// 告诉接收者："不会再发送新数据了"，已经发送的数据还在，可以继续取
	// for range = 一直从通道取数据，取完所有数据后，看到"已关闭"标签，就知道"没有新数据了"，停止等待
	// 如果不 close，for range 会一直等待新数据，导致死锁
	close(queue)
	// 不 close：像在传送带前一直等，不知道是否还有货物
	// close：像在传送带末端贴"已结束"标签，看到后就知道可以停止
	
	// 问题：range 实际上也是一种读取操作吧？
	// 答案：是的！range 在通道上就是循环读取操作
	// for j := range jobs 等价于循环执行 j := <-jobs
	// 每次循环都会阻塞等待，直到有值或通道关闭
	for elem := range queue {
		fmt.Println(elem)
	}
}
