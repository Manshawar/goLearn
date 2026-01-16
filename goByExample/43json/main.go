package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// ========== json.Marshal 的作用 ==========
	// Marshal（序列化/编码）：将 Go 数据结构转换为 JSON 格式的字节数组
	// 函数签名：func Marshal(v interface{}) ([]byte, error)
	// - 参数：任意 Go 数据类型
	// - 返回值：JSON 字节数组和错误信息
	// 用途：将数据转换为 JSON 格式，用于网络传输、文件存储等

	// 示例1：布尔值 → JSON
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB)) // 输出: true

	// 示例2：整数 → JSON
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB)) // 输出: 1

	// 示例3：浮点数 → JSON
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB)) // 输出: 2.34

	// 示例4：字符串 → JSON
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB)) // 输出: "gopher"（带引号）

	// 示例5：切片 → JSON 数组
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB)) // 输出: ["apple","peach","pear"]

	// ========== 结构体序列化 ==========
	// 为什么使用指针 &response1{} 而不是直接 response1{}？
	// 1. 避免复制大结构体，只传递指针（8字节）而不是整个结构体
	// 2. 可以在函数间共享和修改同一个对象
	// 3. 更符合 Go 的常见实践（大结构体通常用指针）
	// 注意：json.Marshal 两种方式都支持，但指针更推荐
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// 使用 json tag 自定义字段名
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// ========== json.Unmarshal 的作用 ==========
	// Unmarshal（反序列化/解码）：将 JSON 字节数组转换为 Go 数据结构
	// 函数签名：func Unmarshal(data []byte, v interface{}) error
	// - 参数1：必须是 []byte 类型（不能直接传 string）
	// - 参数2：指向要填充数据的变量的指针
	// 为什么必须是 []byte？
	// - JSON 数据本质上是字节流，[]byte 是 Go 中表示字节序列的标准类型
	// - 从文件、网络等读取的数据都是 []byte，统一接口
	// - 字符串需要转换：[]byte(str)
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(float64)
	fmt.Println(num)
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// 字符串需要转换为 []byte
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res) // 必须使用 []byte(str) 转换
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// ========== json.NewEncoder 和 Encode ==========
	// NewEncoder：创建一个 JSON 编码器，直接写入到 io.Writer（如文件、网络连接等）
	// Encode：将数据编码为 JSON 并直接写入到指定的 Writer
	// 优势：
	// - 直接写入流，不需要先 Marshal 再写入，更高效
	// - 适合写入文件、网络响应等场景
	// - 自动添加换行符
	// 对比：
	// - json.Marshal：返回 []byte，需要手动写入
	// - json.NewEncoder().Encode：直接写入 Writer，一步完成
	enc := json.NewEncoder(os.Stdout) // 创建编码器，输出到标准输出
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d) // 编码并直接写入 os.Stdout，自动添加换行符
	// 输出: {"apple":5,"lettuce":7}
}
