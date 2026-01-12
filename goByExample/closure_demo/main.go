package main

import "fmt"

// 示例1：经典闭包 - 返回函数（这是你理解的闭包）
func makeCounter() func() int {
	count := 0 // 外部作用域的变量
	return func() int {
		count++ // 捕获并修改外部变量
		return count
	}
}

// 示例2：捕获外部变量的匿名函数（你的代码中的情况）
func closureInGoroutine() {
	x := 10
	go func() {
		// 这个匿名函数捕获了外部变量 x
		// 即使它不是被返回，而是立即执行，它仍然是闭包
		fmt.Println("捕获的外部变量 x =", x)
	}()
}

// 示例3：对比 - 不是闭包的情况
func notClosure() {
	x := 10
	go func(y int) {
		// 这里通过参数传递，没有捕获外部变量
		// 所以这不是闭包（或者说不依赖闭包特性）
		fmt.Println("通过参数传递的 y =", y)
	}(x) // x 作为参数传递，不是捕获
}

func main() {
	fmt.Println("=== 示例1：返回函数的闭包（经典理解）===")
	counter1 := makeCounter()
	counter2 := makeCounter()

	fmt.Println("counter1:", counter1()) // 1
	fmt.Println("counter1:", counter1()) // 2
	fmt.Println("counter2:", counter2()) // 1 (独立的作用域)
	fmt.Println("counter1:", counter1()) // 3

	fmt.Println("\n=== 示例2：在 goroutine 中的闭包（你的代码的情况）===")
	fmt.Println("说明：匿名函数捕获了外部变量，即使不返回也是闭包")

	// 模拟你的代码
	for i := 1; i <= 3; i++ {
		i := i // 局部变量
		go func() {
			// 捕获了外部的 i 变量
			fmt.Printf("Goroutine 捕获的 i = %d\n", i)
		}()
	}

	fmt.Println("\n=== 关键理解 ===")
	fmt.Println("闭包 = 函数 + 捕获的外部变量")
	fmt.Println("闭包不一定要返回函数，只要函数能够访问外部变量就是闭包")
	fmt.Println("在你的代码中，匿名函数捕获了 i 和 wg，所以是闭包")
}
