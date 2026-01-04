package main

import "fmt"

func MapKey[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}
type element[T any] struct {
	value T
	next  *element[T] // 如果不用指针 直接使用element 它会认为是一个深克隆 不断克隆
}
type Stack struct {
	A int
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{value: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{value: v}
		lst.tail = lst.tail.next
	}
}
func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.value)
	}
	return elems
}
func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println(MapKey(m))
	_ = MapKey[int, string](m)

	list := List[int]{}
	list.Push(1)
	list.Push(2)
	list.Push(3)
	fmt.Println(list.GetAll(), list)
	fmt.Printf("head: %v, tail: %v\n", list.head.next.next, list.tail)
}
