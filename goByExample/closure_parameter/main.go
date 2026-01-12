package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Worker %d finished (传入的值是 %d)\n", id, id)
}

func main() {
	fmt.Println("=== 理解闭包和参数传递的区别 ===\n")

	// 情况1：有 i := i（局部变量）的情况
	fmt.Println("情况1：有 i := i（正确的闭包捕获）")
	for i := 1; i <= 3; i++ {
		i := i // 创建局部变量，每次循环都有一个独立的 i
		go func() {
			// 匿名函数捕获了外部的局部变量 i
			// worker(i) 是把捕获到的 i 的值 copy 给 worker 的参数
			worker(i)
		}()
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Println()

	// 情况2：没有 i := i 的情况（经典的闭包陷阱）
	fmt.Println("情况2：没有 i := i（闭包陷阱 - 所有 goroutine 共享同一个 i）")
	for i := 1; i <= 3; i++ {
		go func() {
			// 匿名函数捕获了外部的循环变量 i（所有 goroutine 共享）
			// 当 goroutine 执行时，循环可能已经结束，i 变成了 4
			worker(i) // 这里传入的 i 可能是 4（最终值）
		}()
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Println()

	// 情况3：通过参数传递（不依赖闭包特性）
	fmt.Println("情况3：通过参数传递（不是闭包问题）")
	for i := 1; i <= 3; i++ {
		go func(id int) {
			// id 是参数，每次调用时 i 的值被 copy 到 id
			// 不需要捕获外部变量，所以不会有闭包陷阱
			worker(id)
		}(i) // i 的值在这里被 copy 到参数 id
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Println()

	// 演示闭包捕获变量的本质
	fmt.Println("=== 闭包的本质演示 ===")
	x := 10
	go func() {
		// 这个匿名函数捕获了外部的 x
		fmt.Printf("闭包内部访问外部变量 x = %d\n", x)
		// 注意：这里访问的是外部变量 x，不是 copy
		// 如果在这里修改 x（如果是引用类型），会影响外部的 x
	}()
	time.Sleep(100 * time.Millisecond)

	// 说明闭包和参数传递的区别
	fmt.Println("\n=== 关键区别 ===")
	fmt.Println("1. worker(i) 确实是值传递（copy）- 你说得对！")
	fmt.Println("2. 但是闭包指的是：匿名函数能够访问外部的变量 i")
	fmt.Println("3. 没有 i := i 时，所有 goroutine 捕获的是同一个 i（共享）")
	fmt.Println("4. 有 i := i 时，每个 goroutine 捕获的是独立的局部变量 i")
	fmt.Println("5. 闭包是关于'捕获变量'，不是关于'参数传递'")
}
