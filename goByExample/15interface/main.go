package main

import "fmt"

// 问题：接口和方法的关系是什么？
// 答案：
// 1. 接口定义方法签名的集合（不实现）
// 2. 类型通过实现方法来满足接口（隐式实现，不需要显式声明）
// 3. 接口让不同类型可以统一处理（多态）
// 方法签名（Method Signature）= 单个方法的签名（方法名、参数类型、返回值类型）
// 接口（Interface）= 方法签名的集合
type typeFoo interface {
	foo() string  // 这是方法签名，只有签名，没有实现
}

type foo struct {
	name string
}

// foo 实现了 typeFoo 接口（因为它有 foo() string 方法）
// 问题：它的接口是 typeFoo 吗？
// 答案：是的，foo 实现了 typeFoo 接口，或者说 typeFoo 是 foo 实现的接口
func (f *foo) foo() string {
	return f.name
}

func main() {
	f := foo{name: "foo"}
	// 接口变量可以存储实现类型
	var t typeFoo = &f  // typeFoo 是接口类型，存储了 foo 类型的值
	fmt.Println(t.foo())
}
