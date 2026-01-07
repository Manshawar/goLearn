package main

import "fmt"

type person struct {
	name  string
	age   int
	email string
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}
func main() {
	p := person{name: "John", age: 30, email: "john@example.com"}
	fmt.Println(p)
	fmt.Println(person{name: "Jane", age: 25})
	fmt.Println(newPerson("John"))
	s := newPerson("John")
	sp := &s
	// 问题：作为指针返回的话，一定要解引用吗？可以直接处理吗？
	// 答案：对于单层指针（*person），可以直接使用 . 访问字段，Go 会自动解引用
	// 不需要显式使用 * 来解引用
	// 对于双重指针（**person），需要先解引用一次，得到 *person 后再访问
	(*sp).age = 51  // 双重指针需要解引用一次
	s.age = 52      // 单层指针可以直接访问
	sp2 := &p
	sp2.age = 51    // ✅ 单层指针可以直接访问，不需要 (*sp2).age
	fmt.Println((*sp).age)
	fmt.Println(sp2, *sp)
}
