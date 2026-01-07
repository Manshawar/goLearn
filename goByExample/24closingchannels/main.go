package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	// 问题：为什么 for 循环里直接可以套 if？这个 for 循环是什么用法？
	// 答案：
	// 1. for { } 是无限循环（等价于 while(true)），直到遇到 break 或 return
	// 2. 循环里使用 if 是标准写法，用于条件判断
	// 3. 通道接收返回两个值：值和布尔值（表示通道是否已关闭）
	go func() {
		for {  // 无限循环
			j, more := <-jobs  // 从通道接收值，more 表示通道是否还有值（是否已关闭）
			if more {  // 如果通道未关闭
				fmt.Println("received job", j)
			} else {  // 如果通道已关闭
				fmt.Println("received all jobs")
				done <- true
				return  // 退出循环（退出函数）
			}
		}
	}()
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
