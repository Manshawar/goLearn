package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}
func main() {
	coffee := &Plant{Id: 27, Name: "Coffee", Origin: []string{"Ethiopia", "Brazil"}}
	coffee.Origin = append(coffee.Origin, "Colombia")
	fmt.Println(coffee)
	bulbs := &Plant{Id: 99, Name: "Bulbs", Origin: []string{"Netherlands", "Bulgaria"}}
	fmt.Println(bulbs)
	out, _ := xml.MarshalIndent(bulbs, " ", "  ")
	// fmt.Println(string(out))
	fmt.Println(xml.Header + string(out))
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// ========== 为什么使用 []*Plant（指针切片）而不是 []Plant（值切片）？==========
	// 1. 性能优势：避免复制大结构体
	//    - []Plant：每个元素都是完整的 Plant 结构体副本
	//    - []*Plant：每个元素只是指针（8字节），指向实际的 Plant 对象
	// 2. 共享对象：多个地方可以引用同一个 Plant 对象
	//    - 修改一个引用，所有引用都看到变化
	// 3. nil 值：可以表示"不存在"的元素
	//    - []*Plant 可以包含 nil，表示缺失的元素
	//    - []Plant 不能表示"不存在"（零值不代表不存在）
	// 4. XML 解析灵活性：xml.Unmarshal 处理指针切片更灵活
	//    - 可以动态分配内存
	//    - 可以跳过某些元素（设为 nil）
	// 5. 修改元素：可以直接修改切片中的元素
	//    - []*Plant：修改指针指向的对象
	//    - []Plant：修改的是副本，不影响原对象
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"plant"` // 使用指针切片，更高效和灵活
		// 注意：这里 []*Plant 是类型声明，不是初始化，所以不需要 &
		// *Plant 表示"指向 Plant 的指针类型"，这是类型定义的一部分
	}

	// ========== "指针"的概念 ==========
	// "指针"指的是：类型为 *T 的变量，它存储的是另一个变量的地址
	//
	// 关键理解：
	// - "指针"不是 & 也不是 *，而是类型为 *T 的变量本身
	// - * 在类型声明中表示"指针类型"（如 *Plant）
	// - & 是获取地址的操作符，用于创建指针值
	// - 变量 p（如果类型是 *Plant）就是指针
	//
	// 示例：
	//   var p *Plant        // p 就是指针（类型是 *Plant）
	//   p = &Plant{...}     // & 获取地址，赋值给指针 p
	//   value := *p          // * 解引用，获取指针指向的值
	//
	// ========== * 和 & 的名称和作用 ==========
	// * 的名称：
	//   - 在类型声明中：叫"指针类型声明符"或"指针类型标记"
	//   - 在值使用时：叫"解引用操作符"或"间接引用操作符"
	// & 的名称：
	//   - 叫"取地址操作符"或"地址操作符"
	//
	// 记忆方法：
	// - 类型引用（声明）使用 *：*Plant 表示"指向 Plant 的指针类型"
	// - 实例化（创建值）使用 &：&Plant{} 表示"获取 Plant 值的地址，得到指针"
	// - "指针"是变量本身（类型为 *T），不是符号
	//
	// ========== 什么时候需要 &，什么时候不需要？==========
	// 1. 类型声明：使用 *（不需要 &）
	//    []*Plant 表示"指针切片类型"，这是类型定义，不是值
	//    *Plant 中的 * 是类型声明符，表示"指针类型"
	// 2. 初始化切片元素：使用 &（需要 &）
	//    因为切片类型是 []*Plant，所以每个元素必须是指针
	//    &Plant{} 中的 & 是取地址操作符，用于获取值的地址
	nesting := &Nesting{
		Plants: []*Plant{
			{Id: 1, Name: "Apple"},  // ✅ 需要 &：创建指针
			{Id: 2, Name: "Banana"}, // ✅ 需要 &：创建指针
		},
	}

	// 或者先创建指针变量，再放入切片
	p1 := &Plant{Id: 3, Name: "Cherry"}
	p2 := &Plant{Id: 4, Name: "Date"}
	nesting2 := &Nesting{
		Plants: []*Plant{p1, p2}, // ✅ 直接使用指针变量，不需要 &
	}

	_ = nesting
	_ = nesting2
	// defs := &xml.Decoder{}
	// dec := xml.NewDeer(strings.NewReader(string(out)))
	// dec.Strict = falcodse
}
