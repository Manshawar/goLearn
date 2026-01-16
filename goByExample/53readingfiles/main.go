package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// ========== 方式1：一次性读取整个文件 ==========
	dat, err := os.ReadFile("./tmp/defer.txt")
	check(err)
	fmt.Println(string(dat))

	// ========== 方式2：流式读取（Open + Read）==========
	// Open 的作用：打开文件，返回文件对象，用于后续读取、写入、定位等操作
	// 需要手动关闭文件（defer f.Close()）
	f, err := os.Open("./tmp/defer.txt")
	check(err)
	defer f.Close()

	// Read 为什么需要 []byte？
	// - 文件数据本质上是字节流，Read 需要一个缓冲区来存储读取的数据
	// - []byte 就是这个缓冲区，读取后可以转换为字符串：string(b1)
	// Read 返回什么？
	// - n (int): 实际读取的字节数（可能小于缓冲区大小）
	// - err (error): 错误信息（成功时为 nil）
	// 如果字节数大于缓冲区？
	// - Read 最多只读取缓冲区大小的数据，需要多次调用 Read 才能读取完整文件
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	fmt.Printf("n1 = %d %v\n bytes: %v\n", n1, string(b1[:n1]), b1)

	// Seek：定位到文件特定位置（6 字节处）
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 3)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// ========== ReadAtLeast：确保至少读取指定字节数 ==========
	// io.ReadAtLeast(f, b3, 2) 中的 2 是什么意思？
	// - 2 表示最少要读取的字节数
	// - 如果读取的字节数 < 2，返回错误（io.ErrUnexpectedEOF）
	// - 如果读取的字节数 >= 2，返回成功
	// 显示汉字的问题：
	// - 中文字符在 UTF-8 中通常需要 3 字节
	// - 缓冲区要足够大（至少 10-20 字节），确保能完整读取汉字
	// - 使用 %s 格式化输出，确保显示正确
	o3, err := f.Seek(7, 0)
	check(err)
	b3 := make([]byte, 20)              // 缓冲区足够大，确保能完整读取汉字
	n3, err := io.ReadAtLeast(f, b3, 2) // 至少读取 2 字节
	check(err)
	fmt.Printf("%d bytes @ %d", n3, o3)
	fmt.Printf("%s\n", string(b3[:n3])) // 使用 %s 显示汉字

	// ========== Peek：预览数据，不移动文件指针 ==========
	// Peek 的作用：
	// - 预览指定数量的字节，但不移动文件指针
	// - 数据仍然保留在缓冲区中，可以再次读取
	// - 用于"偷看"数据，决定后续如何处理
	// Peek 和 Read 的区别：
	// - Read：读取数据并移动文件指针，数据被"消耗"掉，无法再次读取
	// - Peek：预览数据但不移动文件指针，数据不被"消耗"，可以再次读取
	// 使用场景：
	// - 需要根据数据内容决定如何处理时（如协议解析、格式判断）
	// - 需要"偷看"数据但不影响后续读取时
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5) // 预览 5 字节，文件指针不移动
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	f.Close()
	// 注意：Peek 后，文件指针还在原位置，可以继续用 Read 读取相同的数据
}
