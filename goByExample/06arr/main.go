package main

import "fmt"

func main() {
	// 使用接口类型，这样 a 是普通数组，但 a[4] 可以是多维数组
	var a [5]interface{}

	// 其他元素是普通 int
	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[3] = 3

	// 只有 a[4] 是多维数组（二维切片）
	a[4] = []int{1, 2}

	fmt.Println("a =", a)
	fmt.Println("a[4] =", a[4])
	fmt.Println("a len", len(a), "a[4] len", len(a[4].([]int)))
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b =", b)

	var twoD [2][3]int
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[i]); j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD =", twoD)
}
