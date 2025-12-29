package main

import "fmt"

type typeFoo interface {
	foo() string
}
type foo struct {
	name string
}

func (f *foo) foo() string {
	return f.name
}
func main() {
	f := foo{name: "foo"}
	var t typeFoo = &f
	fmt.Println(t.foo())
}
