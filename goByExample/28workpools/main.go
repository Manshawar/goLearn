package main

import (
	"fmt"
	"time"
)

// worker 函数：处理任务的工人
// id: worker 的编号
// jobs <-chan int: 只接收通道，从任务队列接收任务
// result chan<- int: 只发送通道，向结果队列发送结果
func worker(id int, jobs <-chan int, result chan<- int) {
	// range 实际上是一种读取操作，等价于循环执行 j := <-jobs
	// 关键点1：range 在通道上会阻塞等待，直到有值或通道关闭
	// 关键点2：channel 循环和 int 循环的区别：
	//   - int 循环（for i := 0; i < 5; i++）：每个 worker 都执行 5 次，总共 15 次
	//   - channel 循环（for j := range jobs）：所有 worker 总共执行 5 次（通道中有 5 个值）
	//     不是每个 worker 执行 5 次，而是所有 worker 共同处理 5 个任务
	// 关键点3：当 jobs <- j 写入任务时，会"唤醒"等待的 worker，worker 开始处理
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		result <- j * 2
	}
	// 当 jobs 通道关闭且所有值被接收后，range 循环自动退出
}

func main() {
	const numJobs = 5
	// 创建任务队列（缓冲通道，容量为 5）
	jobs := make(chan int, numJobs)
	// 创建结果队列（缓冲通道，容量为 5）
	results := make(chan int, numJobs)

	// 启动 3 个 worker goroutine
	// 注意：worker 启动后立即进入 for range 循环，但此时 jobs 通道是空的
	// 所以 worker 会阻塞等待，直到有任务写入
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送 5 个任务到任务队列
	// 关键点：当 jobs <- j 写入时，会"唤醒"等待的 worker
	// worker 的工作时间：从收到任务（jobs <- j）开始，到发送结果（result <- j*2）结束
	for j := 1; j <= numJobs; j++ {
		jobs <- j // 写入任务，唤醒等待的 worker
	}

	// close(jobs) 的作用：
	// 1. 告诉所有 worker：没有更多任务了
	// 2. 让 worker 的 for range 循环在接收完所有任务后自动退出
	// 3. 如果不 close，worker 的 range 会一直等待新任务，导致 goroutine 无法退出
	close(jobs)

	// result 读取操作的作用：
	// 1. 等待所有 worker 完成：接收 5 个结果，确保所有任务都处理完
	// 2. 防止主程序过早退出：如果主程序立即退出，worker 可能还没完成就被终止
	// 关键点：channel 读取（<-results）会产生等待，会阻塞直到有值
	for a := 1; a <= numJobs; a++ {
		<-results // 阻塞等待，直到收到结果
	}
}
