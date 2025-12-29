package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}
func zeroptr(iptr *int) {
	*iptr = 0
}
func main() {
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
	type fn func(int) int
	fn2 := func(n int) fn {
		return func(n int) int {
			return n + 1
		}
	}
	fmt.Println("fn:", fn2(1))
}
