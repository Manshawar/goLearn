package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	// 问题：time.Sleep(time.Second) 有什么用？
	// 答案：让当前 goroutine 暂停执行 1 秒
	// 作用：防止主程序立即退出，给 goroutine 时间执行
	// 如果 main 函数立即结束，goroutine 可能还没执行完就被终止
	time.Sleep(time.Second)
	fmt.Println("done")

}
