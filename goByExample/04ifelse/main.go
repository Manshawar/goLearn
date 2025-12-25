package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 even")
	} else {
		fmt.Println("7 odd")
	}
	if 8%4 == 0 {
		fmt.Println("8 % 4")
	}
	if num := 9; num < 0 {
		fmt.Println('-')
	} else if num < 10 {
		fmt.Println("9")
	} else {
		fmt.Println(num, "other")
	}
}
