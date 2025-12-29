package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	for i, num := range nums {
		fmt.Println("index:", i, "value:", num)
	}
	fmt.Println("sum:", sum)

	maps := map[string]string{"a": "apple", "b": "banana", "c": "cherry"}
	for k, v := range maps {
		fmt.Println("key:", k, "value:", v)
	}
	for k := range maps {
		fmt.Println("key:", k)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
