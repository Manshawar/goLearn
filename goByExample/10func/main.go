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
func intSeq() func() int {
	i := 0

	res := func() int {
		i++
		return i
	}
	return res
}
func recursion() {
	var fact func(n int) int
	fact = func(n int) int {
		if n == 0 {
			return 1
		}
		return n * fact(n-1)
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
