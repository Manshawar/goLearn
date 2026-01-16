package main

import (
	"fmt"
	"os"
)

func main() {
	panic("a problem")

	// 方法1：直接使用相对路径（在当前目录下创建）
	_, err := os.Create("file")
	if err != nil {
		panic(err)
	}
	fmt.Println("文件已创建在当前目录: file")

	// 方法2：获取当前目录的绝对路径，然后创建文件
	// currentDir, _ := os.Getwd()
	// filePath := filepath.Join(currentDir, "file")
	// _, err := os.Create(filePath)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("文件已创建: %s\n", filePath)
}
