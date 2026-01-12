package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// ========== 设计模式：状态管理协程（Stateful Goroutines）==========
// 核心思想：通过 Channel 通信来管理共享状态，替代 Mutex 锁
// Go 并发哲学："不要通过共享内存来通信，而要通过通信来共享内存"
//
// 这个模式的优势：
// 1. 只有一个协程能直接访问共享数据，避免数据竞争
// 2. 通过 Channel 通信，而不是锁，更符合 Go 的并发模型
// 3. 所有对共享数据的访问都是串行化的，保证线程安全

// 读操作请求结构
type readOp struct {
	key  int
	resp chan int // 用于接收读操作的结果
}

// 写操作请求结构
type writeOp struct {
	key  int
	val  int
	resp chan bool // 用于接收写操作的确认
}

func main() {
	// 统计计数器：用于记录读操作和写操作的次数
	// 注意：这些计数器会被多个协程同时修改，所以需要用原子操作保护
	var readOps uint64
	var writeOps uint64

	// 创建请求 channel：用于接收读/写操作请求
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// ========== 核心：状态管理协程 ==========
	// 这个协程是唯一能直接访问 state map 的协程
	// 其他所有协程必须通过 channel 发送请求来访问 state
	// 这相当于用 Channel 替代了 Mutex 的作用
	go func() {
		// state 是私有数据，只有这个协程能直接访问
		// 这避免了多个协程同时访问导致的数据竞争
		var state = make(map[int]int)

		// 无限循环：持续处理读/写请求
		// select 会阻塞等待，直到有请求到达
		for {
			select {
			case read := <-reads:
				// 处理读请求：从 state 读取值，通过 resp channel 返回
				read.resp <- state[read.key]
			case write := <-writes:
				// 处理写请求：修改 state，通过 resp channel 返回确认
				state[write.key] = write.val
				write.resp <- true
			}
			// 注意：select 会阻塞等待，当两个 channel 都没有数据时会阻塞
			// 这是正常的，协程在等待请求时不会消耗 CPU
		}
	}()

	// ========== 读协程：100个协程并发发送读请求 ==========
	for r := 0; r < 100; r++ {
		go func() {
			for {
				// 创建读请求
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int), // 创建响应 channel
				}
				// 发送读请求到 reads channel
				// 这会阻塞，直到状态管理协程处理这个请求
				reads <- read
				// 等待状态管理协程返回结果
				<-read.resp

				// 统计读操作次数
				// 为什么用原子操作？
				// - 因为 100 个读协程都在同时修改 readOps
				// - 需要原子操作避免数据竞争
				// - 这不是用来"隔离协程"，而是用来"统计计数"
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// ========== 写协程：10个协程并发发送写请求 ==========
	for w := 0; w < 10; w++ {
		go func() {
			for {
				// 创建写请求
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool), // 创建响应 channel
				}
				// 发送写请求到 writes channel
				// 这会阻塞，直到状态管理协程处理这个请求
				writes <- write
				// 等待状态管理协程返回确认
				<-write.resp

				// 统计写操作次数
				// 为什么用原子操作？
				// - 因为 10 个写协程都在同时修改 writeOps
				// - 需要原子操作避免数据竞争
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 等待 1 秒，让协程执行操作
	time.Sleep(time.Second)

	// 读取最终的统计结果（使用原子操作保证读取安全）
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// ========== 总结 ==========
	// 1. Channel 的作用：管理状态访问（替代 Mutex）
	//    - 不是用来"记录操作"
	//    - 而是用来"管理共享状态的访问"
	//    - 只有一个协程能直接访问 state，其他协程通过 channel 请求
	//
	// 2. 原子操作的作用：统计计数（避免数据竞争）
	//    - 不是用来"隔离循环协程"
	//    - 而是用来"统计操作次数"
	//    - 因为多个协程都在修改同一个计数器，需要原子操作
	//
	// 3. 这个模式实现了类似 Mutex 的效果，但通过通信而非锁
	//    - Mutex：多个协程直接访问共享数据，用锁保护
	//    - Channel：只有一个协程访问数据，其他协程通过 channel 通信
	//
	// 4. 这是 Go 并发哲学的体现：
	//    "不要通过共享内存来通信，而要通过通信来共享内存"
}
