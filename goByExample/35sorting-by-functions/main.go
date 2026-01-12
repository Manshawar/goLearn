package main

import (
	"fmt"
	"sort"
	"time"
)

// ========== 示例1：按字符串长度排序 ==========
type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// ========== 示例2：按时间排序 ==========
// 定义包含时间字段的结构体
type Event struct {
	Name      string
	Timestamp time.Time
}

// 按时间排序的类型
type byTime []Event

// 实现 sort.Interface 的三个方法
func (e byTime) Len() int {
	return len(e)
}

func (e byTime) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// Less: 按时间升序排序（早的时间在前）
func (e byTime) Less(i, j int) bool {
	return e[i].Timestamp.Before(e[j].Timestamp)
}

// 按时间降序排序的类型（可选）
type byTimeDesc []Event

func (e byTimeDesc) Len() int {
	return len(e)
}

func (e byTimeDesc) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// Less: 按时间降序排序（晚的时间在前）
func (e byTimeDesc) Less(i, j int) bool {
	return e[i].Timestamp.After(e[j].Timestamp)
}

func main() {
	// 示例1：按字符串长度排序
	fmt.Println("=== 按字符串长度排序 ===")
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

	// 示例2：按时间排序
	fmt.Println("\n=== 按时间排序（升序：从早到晚）===")
	events := []Event{
		{Name: "会议", Timestamp: time.Date(2024, 1, 15, 14, 0, 0, 0, time.UTC)},
		{Name: "早会", Timestamp: time.Date(2024, 1, 15, 9, 0, 0, 0, time.UTC)},
		{Name: "午餐", Timestamp: time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)},
		{Name: "培训", Timestamp: time.Date(2024, 1, 15, 16, 0, 0, 0, time.UTC)},
	}

	// 按时间升序排序
	sort.Sort(byTime(events))
	for _, event := range events {
		fmt.Printf("%s: %s\n", event.Name, event.Timestamp.Format("15:04:05"))
	}

	// 按时间降序排序
	fmt.Println("\n=== 按时间排序（降序：从晚到早）===")
	sort.Sort(byTimeDesc(events))
	for _, event := range events {
		fmt.Printf("%s: %s\n", event.Name, event.Timestamp.Format("15:04:05"))
	}
}
