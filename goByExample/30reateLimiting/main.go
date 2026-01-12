package main

import (
	"fmt"
	"time"
)

func main() {
	// ========== 第一部分：基础速率限制 ==========
	// 速率限制（Rate Limiting）的目的：
	// - 限制处理速度（每秒/每分钟处理多少个请求）
	// - 不是限制并发数量（那是并发限制/信号量的作用）
	// - 例如：每 200ms 处理 1 个请求 = 每秒最多 5 个请求

	request := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		request <- i
	}
	close(request)

	// time.Tick 返回一个 channel，每 200ms 产生一个时间值
	// 这是 Go 中定时器的常见模式：time.Tick + range
	limiter := time.Tick(200 * time.Millisecond)
	for req := range request {
		<-limiter // 等待令牌：每 200ms 才能拿到一个，限制处理速度
		fmt.Println("request", req, time.Now())
	}

	// ========== 第二部分：带突发的速率限制（Bursty Rate Limiting）==========
	// Bursty 的含义：允许前几个请求立即处理（突发），后续请求按速率限制

	// 创建一个带缓冲的 channel，容量为 3
	// 缓冲的作用：允许预填充令牌，实现突发处理
	burstyLimiter := make(chan time.Time, 3)

	// 预填充 3 个令牌，这样前 3 个请求可以立即处理（不需要等待）
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 启动一个协程持续产生令牌
	// 为什么放在协程中？
	// - 因为写入操作可能阻塞（当缓冲满时）
	// - 阻塞操作放在协程中，避免阻塞主流程
	// - 当缓冲满时，写入会阻塞等待，直到有读取释放空间
	go func() {
		// time.Tick + range：每 200ms 产生一个时间值
		// 这是 Go 中实现定时器的标准模式
		for t := range time.Tick(200 * time.Millisecond) {
			// 写入令牌到 channel
			// 关键理解：
			// - 缓冲未满时：写入立即成功，不阻塞
			// - 缓冲已满时：写入会阻塞，等待有空间
			// - 一旦主流程读取（<-burstyLimiter），释放空间，写入继续
			burstyLimiter <- t
			fmt.Println("burstyLimiter", t)
		}
	}()

	// 方式1：直接循环（静态，无法动态添加请求）
	for i := 1; i <= 5; i++ {
		<-burstyLimiter // 等待令牌：前 3 个立即获得，后续每 200ms 一个
		fmt.Println("request", i, time.Now())
	}

	// ========== 第三部分：使用 channel 实现动态添加请求 ==========
	// 为什么要用 channel 而不是直接循环？
	// - 直接循环：只能处理固定数量的请求，无法在运行时动态添加
	// - 使用 channel：可以从多个来源动态添加请求，更灵活
	//   例如：网络请求、用户操作、其他协程等都可以随时添加请求

	burstyRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)

	// 使用 range 遍历 channel，处理所有动态添加的请求
	// 这种模式的优势：
	// - 请求可以随时从多个来源添加（网络、用户、其他协程等）
	// - 主循环统一处理，自动应用速率限制
	// - 当 channel 关闭时，循环自动结束
	for req := range burstyRequest {
		<-burstyLimiter // 等待速率限制令牌
		fmt.Println("request", req, time.Now())
	}

	// ========== 总结 ==========
	// 1. 速率限制：限制处理速度（每 200ms 处理 1 个），不是限制并发数量
	// 2. 缓冲机制：预填充允许突发处理，后续按速率限制
	// 3. 阻塞行为：缓冲满时写入阻塞，读取后继续写入
	// 4. 协程使用：阻塞操作放在协程中，避免阻塞主流程
	// 5. Channel 模式：支持动态添加请求，比直接循环更灵活
	// 6. 定时器模式：time.Tick + range 是 Go 中实现定时器的标准方式
}
