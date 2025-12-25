package main

import "fmt"

type Point struct {
    X, Y int
}
func main(){
	pp := &Point{1, 2}
	p2 := *pp     
	p2.X = 99
	fmt.Printf("%#v",pp)
}