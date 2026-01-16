package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("./defer.txt")
	defer closeFile(f)
	writeFile(f)
}

// * 值引用 减少性能开销 避免创建副本
func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
