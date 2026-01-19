package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	// ========== 第一部分：执行 date 命令（简单方式）==========
	// 步骤1：创建命令对象，准备执行 "date" 命令
	// exec.Command() 返回一个 *exec.Cmd 对象，但此时命令还未执行
	dateCmd := exec.Command("date")

	// 步骤2：执行命令并获取输出
	// Output() 方法会：
	//   - 启动命令
	//   - 等待命令执行完成
	//   - 返回标准输出的字节数组和错误信息
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	// 步骤3：打印命令提示和结果
	fmt.Println("> date")
	fmt.Println(string(dateOut)) // 将字节数组转换为字符串并打印

	// ========== 第二部分：执行 grep 命令（使用管道交互）==========
	// 步骤1：创建 grep 命令，查找包含 "hello" 的行
	// 第一个参数是命令名，后续参数是命令的参数
	grepCmd := exec.Command("grep", "hello")

	// 步骤2：获取命令的标准输入管道
	// StdinPipe() 返回一个 io.WriteCloser，可以向命令的标准输入写入数据
	grepIn, _ := grepCmd.StdinPipe()

	// 步骤3：获取命令的标准输出管道
	// StdoutPipe() 返回一个 io.ReadCloser，可以读取命令的标准输出
	grepOut, _ := grepCmd.StdoutPipe()

	// 步骤4：启动命令（但不等待完成）
	// Start() 启动命令后立即返回，命令在后台运行
	grepCmd.Start()

	// 步骤5：向 grep 命令的标准输入写入数据
	// grep 会从标准输入读取数据，查找包含 "hello" 的行
	grepIn.Write([]byte("hello grep\n"))

	// 步骤6：关闭标准输入管道
	// 告诉 grep 命令输入已经结束，可以开始处理了
	grepIn.Close()

	// 步骤7：读取 grep 命令的输出
	// io.ReadAll() 读取所有输出数据，直到 EOF（文件结束）
	grepBytes, _ := io.ReadAll(grepOut)

	// 步骤8：等待命令执行完成
	// Wait() 会阻塞直到命令执行完成，并返回退出状态
	grepCmd.Wait()

	// 步骤9：打印命令提示和结果
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// ========== 第三部分：通过 bash 执行复杂命令 ==========
	// 步骤1：通过 bash 执行命令
	// "bash", "-c", "ls -a -l -h" 表示：
	//   - 使用 bash shell
	//   - "-c" 表示执行后面的命令字符串
	//   - "ls -a -l -h" 是要执行的命令（列出所有文件，详细信息，人类可读格式）
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")

	// 步骤2：执行命令并获取输出（同第一部分）
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}

	// 步骤3：打印命令提示和结果
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
