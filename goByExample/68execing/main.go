package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// ========== 使用 syscall.Exec 的流程 ==========
	// 步骤1：先查找命令的完整路径（二进制文件位置）
	// LookPath 会在系统的 PATH 环境变量中搜索 "ls" 命令
	// 返回找到的完整路径，例如："/bin/ls" 或 "/usr/bin/ls"
	// 为什么要先查找？
	//   - syscall.Exec 需要完整的二进制文件路径，不能只给命令名
	//   - 确保命令存在，避免执行失败
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// 步骤2：准备命令参数
	// args[0] 必须是程序名（通常与 binary 的最后一个部分相同）
	// 后续元素是传递给程序的参数
	args := []string{"ls", "-a", "-l", "-h"}

	// 步骤3：获取当前进程的环境变量
	// 这些环境变量会被传递给新执行的程序
	env := os.Environ()

	// 步骤4：执行二进制文件（替换当前进程）
	// syscall.Exec 与 exec.Command 的区别：
	//   - exec.Command：创建子进程执行，原进程继续运行
	//   - syscall.Exec：用新程序替换当前进程，原进程代码不会继续执行
	//   如果 Exec 成功，这行之后的代码永远不会执行！
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
	// 注意：如果 Exec 成功，这里的代码不会执行
	// 因为当前进程已经被 ls 命令替换了
}
