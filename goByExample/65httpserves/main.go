package main

import (
	"fmt"
	"net/http"
)

// 注意参数类型：
// - w http.ResponseWriter：接口类型，不需要 *（接口本身就是引用语义）
// - req *http.Request：结构体类型，需要 *（避免复制大结构体，提高性能）
func hello(w http.ResponseWriter, req *http.Request) {
	// fmt.Fprintf 的第一个参数是 io.Writer 接口
	// - 如果传入 os.Stdout，会输出到控制台
	// - 如果传入 http.ResponseWriter（w），会输出到 HTTP 响应（网页）
	// w 是 ResponseWriter，实现了 io.Writer 接口，所以内容会写入到 HTTP 响应中
	fmt.Fprintf(w, "hello\n")
}
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8090", nil)
}
