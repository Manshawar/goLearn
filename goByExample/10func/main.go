package main

import "fmt"

func main() {
	fmt.Println("func:")
	fmt.Println("plus:", plus(1, 2))
	fmt.Println("plusPlus:", plusPlus(1, 2, 3))
	a, b := vals()
	fmt.Println("a:", a, "b:", b)
	_, c := vals()
	fmt.Println("c:", c)
	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(intSeq()())
	fmt.Println("recursion:")
	recursion()
}
func plus(a int, b int) int {
	return a + b
}
func plusPlus(a, b, c int) int {
	return a + b + c
}
func vals() (int, int) {
	return 3, 7
}
func sum(nums ...int) int {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
	return total
}
// 问题：Go 不能像 TypeScript 一样返回其他函数，必须写在 return 中吗？
// 答案：可以！Go 支持先定义函数变量再返回，不一定必须写在 return 中
// 函数是一等公民，可以作为类型使用，可以赋值给变量、作为参数传递、作为返回值
func intSeq() func() int {
	i := 0

	// 方式1：先定义函数变量，再返回（类似 TypeScript）
	res := func() int {
		i++
		return i
	}
	return res
	// 也可以直接在 return 中写：return func() int { i++; return i }
}

// 问题：为什么递归函数在函数内部定义会报错？
// 答案：Go 不支持在函数内部定义命名函数，只能定义匿名函数
// 问题：为什么使用 := 定义递归匿名函数会报错？
// 答案：使用 := 时，变量在赋值完成前不存在，无法在函数体内引用自身
// 必须先用 var 声明，再赋值，这样函数体内可以引用自身
func recursion() {
	// ✅ 正确：先声明变量
	var fact func(n int) int
	// ✅ 再赋值（此时 fact 已经存在，可以在函数体内引用）
	fact = func(n int) int {
		if n == 0 {
			return 1
		}
		return n * fact(n-1)  // 现在可以引用 fact 了
	}
	fmt.Println("fact:", fact(7))

	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println("fib:", fib(7))
}
