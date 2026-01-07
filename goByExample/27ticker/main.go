package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Done")
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	// 问题：为什么打印 "Ticker stopped" 才会触发 "Done"？
	// 答案：打印语句本身不会触发，触发的是 done <- true
	// 打印语句只是给了 goroutine 处理时间，让 "Done" 有机会输出
	// 如果主程序立即退出，goroutine 可能来不及打印 "Done"
	done <- true  // ✅ 这里触发！向 done 通道发送值，触发 case <-done
	// fmt.Println("Ticker stopped")  // 这行给了 goroutine 处理时间
	time.Sleep(100 * time.Millisecond) // 用这个等待，也能看到 "Done"
}
