package main

import (
	"fmt"
	"sort"
)

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	// 构造示例图
	addEdge("A", "B")
	addEdge("A", "C")
	addEdge("B", "C")
	addEdge("C", "D")
	addEdge("D", "A") // 形成一个环
	addEdge("E", "B") // E -> B

	// 方式一：直接打印整个 map（顺序不定）
	fmt.Printf("直接打印 graph: %v\n\n", graph)

	// 方式二：按节点和邻居有序打印（可读性更好）
	fmt.Println("有序打印 graph:")
	// 获取并排序节点名
	nodes := make([]string, 0, len(graph))
	for n := range graph {
		nodes = append(nodes, n)
	}
	sort.Strings(nodes)

	for _, n := range nodes {
		neighbors := make([]string, 0, len(graph[n]))
		for nb := range graph[n] {
			neighbors = append(neighbors, nb)
		}
		sort.Strings(neighbors)
		fmt.Printf("%s -> %v\n", n, neighbors)
	}

	// 示例 hasEdge 用法
	fmt.Println()
	fmt.Printf("hasEdge(A, B) = %v\n", hasEdge("A", "B"))
	fmt.Printf("hasEdge(B, A) = %v\n", hasEdge("B", "A"))
	fmt.Printf("graph: %v\n", graph)
}
