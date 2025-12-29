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
	(*sp).age = 51
	s.age = 52
	sp2 := &p
	sp2.age = 51
	fmt.Println((*sp).age)
	fmt.Println(sp2, *sp)
}
