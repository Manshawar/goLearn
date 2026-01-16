# Go 语言中 * 和 & 的名称和用法

## 名称

### * 的名称
- **在类型声明中**：叫 **"指针类型声明符"** 或 **"指针类型标记"**
- **在值使用时**：叫 **"解引用操作符"** 或 **"间接引用操作符"**

### & 的名称
- 叫 **"取地址操作符"** 或 **"地址操作符"**

## 用法总结

### 记忆方法
- **类型引用（声明）使用 `*`**：`*Plant` 表示"指向 Plant 的指针类型"
- **实例化（创建值）使用 `&`**：`&Plant{}` 表示"获取 Plant 值的地址，得到指针"

## 详细说明

### 1. * 的两种用法

#### 用法1：类型声明中的 *（指针类型声明符）
```go
// * 在类型声明中，表示"指针类型"
var p *Plant        // p 的类型是 *Plant（指向 Plant 的指针）
var plants []*Plant // plants 的类型是 []*Plant（指针切片）

type Nesting struct {
    Plants []*Plant  // *Plant 是类型定义的一部分
}
```

#### 用法2：值使用中的 *（解引用操作符）
```go
var p *Plant = &Plant{Id: 1, Name: "Apple"}
value := *p  // *p 表示"解引用"，获取指针指向的值
// value 的类型是 Plant（值类型）
```

### 2. & 的用法（取地址操作符）

```go
// & 用于获取值的地址，得到指针
p1 := Plant{Id: 1, Name: "Apple"}
p2 := &p1  // &p1 获取 p1 的地址，p2 的类型是 *Plant

// 或者直接创建指针
p3 := &Plant{Id: 2, Name: "Banana"}  // 创建 Plant 值，然后取地址
```

## 完整示例

```go
package main

import "fmt"

type Plant struct {
    Id   int
    Name string
}

func main() {
    // ========== 1. 类型声明：使用 * ==========
    var p *Plant        // * 是类型声明符：p 的类型是 *Plant
    var plants []*Plant // * 是类型声明符：plants 的类型是 []*Plant
    
    // ========== 2. 创建指针值：使用 & ==========
    p = &Plant{Id: 1, Name: "Apple"}  // & 是取地址操作符：获取值的地址
    
    // ========== 3. 解引用：使用 * ==========
    value := *p  // * 是解引用操作符：获取指针指向的值
    fmt.Println(value)  // 输出: {1 Apple}
    
    // ========== 4. 类型声明 vs 值创建 ==========
    type Nesting struct {
        Plants []*Plant  // * 是类型声明符：定义指针类型
    }
    
    nesting := &Nesting{  // & 是取地址操作符：获取 Nesting 值的地址
        Plants: []*Plant{
            &Plant{Id: 1, Name: "Apple"},   // & 是取地址操作符
            &Plant{Id: 2, Name: "Banana"},  // & 是取地址操作符
        },
    }
    
    // ========== 5. 解引用示例 ==========
    if nesting.Plants[0] != nil {
        plant := *nesting.Plants[0]  // * 是解引用操作符：获取指针指向的值
        fmt.Println(plant.Name)  // 输出: Apple
    }
}
```

## 对比表

| 符号 | 名称 | 位置 | 作用 | 示例 |
|------|------|------|------|------|
| `*` | 指针类型声明符 | 类型声明 | 声明指针类型 | `var p *Plant` |
| `*` | 解引用操作符 | 值使用 | 获取指针指向的值 | `value := *p` |
| `&` | 取地址操作符 | 值使用 | 获取值的地址 | `p := &Plant{}` |

## 记忆技巧

1. **类型声明用 `*`**：`*Plant` = "指向 Plant 的指针类型"
2. **创建指针用 `&`**：`&Plant{}` = "获取 Plant 值的地址"
3. **获取值用 `*`**：`*p` = "解引用，获取指针指向的值"

## 常见错误

```go
// ❌ 错误：类型声明中不能使用 &
type Nesting struct {
    Plants []&Plant  // 错误！& 不能用于类型声明
}

// ✅ 正确：类型声明使用 *
type Nesting struct {
    Plants []*Plant  // 正确
}

// ❌ 错误：值初始化中不能使用 *
Plants: []*Plant{
    *Plant{Id: 1},  // 错误！* 不能用于创建值
}

// ✅ 正确：值初始化使用 &
Plants: []*Plant{
    &Plant{Id: 1},  // 正确
}
```
