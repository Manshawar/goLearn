package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	// 示例1：基本排序测试
	values1 := []int{3, 1, 4, 1, 5}
	fmt.Printf("排序前: %v\n", values1)
	Sort(values1) // 在这里打断点，观察 values1 的变化
	fmt.Printf("排序后: %v\n", values1)
	fmt.Println()

	// // 示例2：逆序数组
	// values2 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	// fmt.Printf("排序前: %v\n", values2)
	// Sort(values2) // 在这里打断点
	// fmt.Printf("排序后: %v\n", values2)
	// fmt.Println()

	// // 示例3：已排序数组
	// values3 := []int{1, 2, 3, 4, 5}
	// fmt.Printf("排序前: %v\n", values3)
	// Sort(values3) // 在这里打断点
	// fmt.Printf("排序后: %v\n", values3)
	// fmt.Println()

	// // 示例4：包含重复元素
	// values4 := []int{5, 2, 8, 2, 5, 1, 8, 1}
	// fmt.Printf("排序前: %v\n", values4)
	// Sort(values4) // 在这里打断点
	// fmt.Printf("排序后: %v\n", values4)
}
