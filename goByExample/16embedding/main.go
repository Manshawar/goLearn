package main

import "fmt"

type base struct {
	num int
}

type container struct {
	base
	str string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num: %v", b.num)
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "hello",
	}
	fmt.Println(co.describe())

	type describe interface {
		describe() string
	}
	var d describe = co
	fmt.Println(d.describe(), co.describe())
}
