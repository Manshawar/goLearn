package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	r.width = 100
	return r.width * r.height
}
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}
func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area(), r)
	fmt.Println("perim:", r.perim(), r)

}
